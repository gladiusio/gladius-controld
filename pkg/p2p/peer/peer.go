package peer

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/gladiusio/gladius-controld/pkg/p2p/state"
	"github.com/hashicorp/memberlist"
	uuid "github.com/satori/go.uuid"
)

// New returns a new peer type
func New() *Peer {
	d := &delegate{}
	hostname, _ := os.Hostname()

	c := memberlist.DefaultWANConfig()
	c.PushPullInterval = 15 * time.Second
	c.GossipInterval = 300 * time.Millisecond
	c.Delegate = d
	c.Name = hostname + "-" + uuid.NewV4().String()

	m, err := memberlist.Create(c)
	if err != nil {
		panic(err)
	}

	queue := &memberlist.TransmitLimitedQueue{
		RetransmitMult: 3,
	}

	peer := &Peer{peerState: &state.State{}, running: false, peerDelegate: d, member: m, PeerQueue: queue}

	queue.NumNodes = func() int { return peer.member.NumMembers() }
	d.peer = peer
	return peer
}

// Peer is a type that represents a peer in the Gladius p2p network.
type Peer struct {
	peerDelegate *delegate
	PeerQueue    *memberlist.TransmitLimitedQueue
	peerState    *state.State
	member       *memberlist.Memberlist
	running      bool
}

type broadcast struct {
	msg    []byte
	notify chan<- struct{}
}

type update struct {
	Action string          // merge
	Data   json.RawMessage // crdt.GCounterJSON
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
// and returns a list of links for where the missing content can be found along
// with their hash
func (p *Peer) CompareContent(contentList []string) []string {
	return []string{}
}
