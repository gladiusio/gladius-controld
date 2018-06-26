package peer

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"time"

	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/gladiusio/gladius-controld/pkg/p2p/state"
)

// New returns a new peer type
func New() *Peer {
	peer := &Peer{peerState: &state.State{}, running: false, maxMessageAge: 1000, client: &client{}}
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
	return errors.New("message signature too old")
}

func (p Peer) pushStateMessage(sm *signature.SignedMessage) error {
	ipList := p.peerState.GetNodeFields("IPAddress")
	if len(ipList) > 0 {
		go func() {
			s := rand.NewSource(time.Now().Unix())
			r := rand.New(s) // initialize local pseudorandom generator
			timestamp := sm.GetTimestamp()
			for (time.Now().Unix() - timestamp) < p.maxMessageAge {
				if len(ipList) > 0 {
					index := r.Intn(len(ipList))
					ipInterface := ipList[index]
					// Delete it for this run
					ipList = append(ipList[:index], ipList[index+1:]...)
					if ipInterface != nil {
						// Get the data from the signed field
						ip := ipInterface.(state.SignedField).Data

						conn, err := net.Dial("tcp", ip+":4351")
						if err != nil {
							fmt.Println("dialing:", err)
						} else {
							client := rpc.NewClient(conn)
							var reply string
							err = client.Call("State.Update", sm, &reply)
							if err != nil {
								fmt.Println("can't call method:", err)
							}
							err = conn.Close()
							if err != nil {
								fmt.Println("can't close connection:", err)
							}
						}

					}
				}
				time.Sleep(100 * time.Millisecond)
			}
		}()
		return nil
	}
	return errors.New("no peers")
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
