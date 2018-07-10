package peer

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"strconv"
	"time"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/p2p/message"

	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/gladiusio/gladius-controld/pkg/p2p/state"
)

// New returns a new peer type
func New() *Peer {
	peer := &Peer{peerState: &state.State{}, running: false, maxMessageAge: 10, client: &client{}}
	peer.server = newServer(peer)
	return peer
}

// Peer is a type that represents a peer in the Gladius p2p network.
type Peer struct {
	peerState     *state.State
	running       bool
	server        *server
	client        *client
	maxMessageAge int64
}

// Start starts the peer
func (p *Peer) Start() {
	p.server.Start()
}

// Stop stops the peer
func (p *Peer) Stop() {

}

// PullState pulls the state from a peer and verifies it before loading it into
// its own state
func (p *Peer) PullState(ip, passphrase string) error {
	currTime := strconv.FormatUint(uint64(time.Now().Unix()), 10)
	m := message.New([]byte("{\"challenge_time\":" + currTime + "}"))
	smString, err := signature.CreateSignedMessageString(m, passphrase)
	sm := &signature.SignedMessage{}
	json.Unmarshal([]byte(smString), sm)
	if err != nil {
		return errors.New("cannot make signed message: " + err.Error())
	}
	client, err := rpc.DialHTTP("tcp", ip+":4351")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("State.Get", sm, &reply)
	if err != nil {
		fmt.Println(reply)
		return errors.New("can't call method: " + err.Error())
	}
	// Convert the incoming json to a State type
	incomingState, err := state.ParseNetworkState([]byte(reply))
	if err != nil {
		fmt.Println(reply)
		return errors.New("corrupted state: " + err.Error())
	}
	// Get the signatures and rebuild the state
	sigList := incomingState.GetSignatureList()
	for _, sig := range sigList {
		p.GetState().UpdateState(sig)
	}

	return nil
}

func (p *Peer) getPeerIPs() []string {
	ipList := p.peerState.GetNodeFields("IPAddress")
	ips := make([]string, 0)
	ga := blockchain.NewGladiusAccountManager()
	address, err := ga.GetAccountAddress()
	myIP := ""
	if err == nil {
		myIP = p.GetState().GetNodeField(address.String(), "IPAddress").(state.SignedField).Data
		fmt.Println("MyIP: " + myIP)
	}
	// Go through all of the fields and get the string IP
	for _, ip := range ipList {
		if ip != myIP {
			ips = append(ips, ip.(state.SignedField).Data)
		}
	}

	return ips
}

// UpdateAndPushState updates the local state and pushes it to several other peers
func (p *Peer) UpdateAndPushState(sm *signature.SignedMessage) error {
	if sm.GetAgeInSeconds() < p.maxMessageAge {
		err := p.peerState.UpdateState(sm)
		if err != nil {
			return err
		}
		// Send to peers
		err = p.pushStateMessage(sm)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("message signature too old: was " + strconv.Itoa(int(sm.GetAgeInSeconds())))
}

func (p Peer) pushStateMessage(sm *signature.SignedMessage) error {
	ipList := p.getPeerIPs()
	numOfPeers := len(ipList)

	if numOfPeers > 0 {
		// Calculate the frequency based on the number of peers to not overload
		// small networks
		// // waitTime := calcWaitTimeMillis(numOfPeers)
		go func() {
			s := rand.NewSource(time.Now().Unix())
			r := rand.New(s) // initialize local pseudorandom generator
			timestamp := sm.GetTimestamp()
			count := 0
			for (time.Now().Unix() - timestamp) < 3 {
				// If we decide to modify the peer list this is useful
				if len(ipList) > 0 {
					index := r.Intn(len(ipList))
					// Get the data from the signed field
					ip := ipList[index]
					client, err := rpc.DialHTTP("tcp", ip+":4351")
					if err != nil {
						fmt.Println("dialing:", err)
					} else {
						var reply string
						err = client.Call("State.Update", sm, &reply)
						if err != nil {
							fmt.Println("can't call method State.Update:", err)
							break
						}
					}
				} else {
					break
				}
				count++
				time.Sleep(10 * time.Millisecond)
			}
		}()
		return nil
	}
	// No data has been sent to peers because there are none
	return errors.New("not enough peers, only updating local state")
}

func calcWaitTimeMillis(peers int) time.Duration {
	if peers > 1000 {
		return 100
	} else if peers > 200 {
		return 200
	} else if peers > 10 {
		return 300
	}
	return 500

}

// GetState returns the current local state
func (p *Peer) GetState() *state.State {
	return p.peerState
}

// CompareContent compares the content provided with the content in the state
// and returns a list of links for where the missing content can be found along
// with their hash
func (p *Peer) CompareContent(contentList []string) []string {
	return []string{}
}
