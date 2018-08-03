package peer

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/deckarep/golang-set"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/gladiusio/gladius-controld/pkg/p2p/state"
	"github.com/hashicorp/memberlist"
	uuid "github.com/satori/go.uuid"
)

// New returns a new peer type
func New(ga *blockchain.GladiusAccountManager) *Peer {
	d := &delegate{}
	md := &mergeDelegate{}
	hostname, _ := os.Hostname()

	c := memberlist.DefaultWANConfig()
	c.PushPullInterval = 15 * time.Second
	c.GossipInterval = 300 * time.Millisecond
	c.ProbeTimeout = 4 * time.Second
	c.ProbeInterval = 8 * time.Second
	c.GossipNodes = 3
	c.Delegate = d
	c.Merge = md
	c.Name = hostname + "-" + uuid.NewV4().String()

	m, err := memberlist.Create(c)
	if err != nil {
		panic(err)
	}

	queue := &memberlist.TransmitLimitedQueue{
		RetransmitMult: 3,
	}

	peer := &Peer{
		peerState:           &state.State{},
		running:             false,
		peerDelegate:        d,
		member:              m,
		PeerQueue:           queue,
		challengeReceiveMap: make(map[string]chan *signature.SignedMessage),
		ga:                  ga,
	}

	queue.NumNodes = func() int { return peer.member.NumMembers() }
	d.peer = peer
	md.peer = peer
	return peer
}

// Peer is a type that represents a peer in the Gladius p2p network.
type Peer struct {
	ga                  *blockchain.GladiusAccountManager
	peerDelegate        *delegate
	PeerQueue           *memberlist.TransmitLimitedQueue
	peerState           *state.State
	member              *memberlist.Memberlist
	running             bool
	challengeReceiveMap map[string]chan *signature.SignedMessage // Map of challenge set ids to a receive channel of the responses from the questioned peers.
	mux                 sync.Mutex
}

type broadcast struct {
	msg    []byte
	notify chan<- struct{}
}

type update struct {
	From   memberlist.Node
	Action string          // Can be "merge", "challenge_question", or "challenge_response"
	Data   json.RawMessage // Usually a signed message, but can also be a challenge question
}

// Used to send to a node through an "update"
type challenge struct {
	ChallengeID string
	Question    string
}

func (b *broadcast) Invalidates(other memberlist.Broadcast) bool {
	return false
}

func (b *broadcast) Message() []byte {
	return b.msg
}

func (b *broadcast) Finished() {
	if b.notify != nil {
		close(b.notify)
	}
}

// Join will request to join the network from a specific node
func (p *Peer) Join(ipList []string) error {
	_, err := p.member.Join(ipList)
	if err != nil {
		return err
	}

	node := p.member.LocalNode()
	fmt.Printf("Local member %s:%d\n", node.Addr, node.Port)

	return nil
}

// StopAndLeave will infomr the network of it leaving and shutdown
func (p *Peer) StopAndLeave() error {
	err := p.member.Leave(1 * time.Millisecond)
	if err != nil {
		return err
	}

	err = p.member.Shutdown()
	if err != nil {
		return err
	}

	return nil
}

func (p *Peer) registerOutgoingChallenge(challengeID string) {
	p.mux.Lock()
	p.challengeReceiveMap[challengeID] = make(chan *signature.SignedMessage)
	p.mux.Unlock()
}

func (p *Peer) getChallengeResponseChannel(challengeID string) (chan *signature.SignedMessage, error) {
	p.mux.Lock()
	defer p.mux.Unlock()
	if challengeChan, ok := p.challengeReceiveMap[challengeID]; ok {
		return challengeChan, nil
	}
	return nil, errors.New("Could not find channel")
}

// UpdateAndPushState updates the local state and pushes it to several other peers
func (p *Peer) UpdateAndPushState(sm *signature.SignedMessage) error {
	err := p.GetState().UpdateState(sm)
	if err != nil {
		return err
	}

	signedBytes, err := json.Marshal(sm)
	if err != nil {
		return err
	}

	b, err := json.Marshal(&update{
		Action: "merge",
		Data:   signedBytes,
		From:   *p.member.LocalNode(),
	})

	if err != nil {
		return err
	}

	p.PeerQueue.QueueBroadcast(&broadcast{
		msg:    b,
		notify: nil,
	})

	return nil
}

// GetState returns the current local state
func (p *Peer) GetState() *state.State {
	return p.peerState
}

// CompareContent compares the content provided with the content in the state
// and returns a list of the missing files names in the format of:
// website/<"asset" or "route">/filename
func (p *Peer) CompareContent(contentList []interface{}) []interface{} {
	contentWeHaveSet := mapset.NewSetFromSlice(contentList)

	contentField := p.GetState().GetPoolField("RequiredContent")
	if contentField == nil {
		return make([]interface{}, 0)
	}
	contentFromPool := contentField.(state.SignedList).Data

	// Convert to an interface array
	s := make([]interface{}, len(contentFromPool))
	for i, v := range contentFromPool {
		s[i] = v
	}

	// Create a set
	contentWeNeed := mapset.NewSetFromSlice(s)

	// Return the difference of the two
	return contentWeNeed.Difference(contentWeHaveSet).ToSlice()
}

// GetContentLinks get's a link for each item in the contentList from a random
// node in the network that has that content
func (p *Peer) GetContentLinks(contentList []interface{}) []string {
	return []string{}
}
