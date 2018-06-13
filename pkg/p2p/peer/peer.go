package peer

import (
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/gladiusio/gladius-controld/pkg/p2p/state"
)

// New returns a new peer type
func New() *Peer {
	return &Peer{peerState: &state.State{}, running: false}
}

// Peer is a type that represents a peer in the Gladius p2p network.
type Peer struct {
	peerState *state.State
	running   bool
	server    *server
	client    *client
}

// Start starts the peer
func (p *Peer) Start() {

}

// Stop stops the peer
func (p *Peer) Stop() {

}

// UpdateAndPushState updates the local state and pushes it to several other peers
func (p *Peer) UpdateAndPushState(sm *signature.SignedMessage) {
	maxMessageAge := int64(10)
	if sm.GetAgeInSeconds() < maxMessageAge {
		p.peerState.UpdateState(sm)
		// Push Message to peers every
	}
}

func (p Peer) pushStateMessage(sm *signature.SignedMessage) {

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
