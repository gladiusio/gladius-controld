package peer

import (
	"encoding/json"

	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/gladiusio/gladius-controld/pkg/p2p/state"
)

type delegate struct {
	peer *Peer
}

func (d *delegate) NodeMeta(limit int) []byte {
	return []byte{}
}

// NotifyMsg is called when a new message is recieved by this peer
func (d *delegate) NotifyMsg(b []byte) {
	var update *update
	if err := json.Unmarshal(b, &update); err != nil {
		panic(err)
	}
	switch update.Action {
	case "merge": // This is when a node is propigating a new message via gossip
		var sm *signature.SignedMessage

		err := json.Unmarshal([]byte(update.Data), &sm)
		if err != nil {
			panic(err)
		}
		go d.peer.GetState().UpdateState(sm)
	case "challenge_response": // This is from a node responding to a challenge question
		var sm *signature.SignedMessage

		err := json.Unmarshal([]byte(update.Data), &sm)
		if err != nil {
			panic(err)
		}

		cBytes, err := sm.Message.MarshalJSON()
		if err != nil {
			panic(err)
		}
		var c *challenge

		err = json.Unmarshal(cBytes, c)
		if err != nil {
			panic(err)
		}

		// Queue up our incomming challenge
		channel, err := d.peer.getChallengeResponseChannel(c.challengeID)
		if err != nil {
			return
		}

		channel <- sm
	case "challenge_question": // This is a node recieving a question
		var c challenge
		err := json.Unmarshal([]byte(update.Data), &c)
		if err != nil {
			panic(err)
		}

	default:
		panic("unsupported update action")
	}

}

// GetBroadcasts returns the list of broadcast messages (not for us)
func (d *delegate) GetBroadcasts(overhead, limit int) [][]byte {
	return d.peer.PeerQueue.GetBroadcasts(overhead, limit)
}

// Get the local state that we can pass on to another node for replication
func (d *delegate) LocalState(join bool) []byte {
	b, err := d.peer.GetState().GetJSON()

	if err != nil {
		panic(err)
	}

	return b
}

// MergeRemoteState is invoked after a TCP Push/Pull. This is the
// state received from the remote side and is the result of the
// remote side's LocalState call. The 'join'
// boolean indicates this is for a join instead of a push/pull.
func (d *delegate) MergeRemoteState(buf []byte, join bool) {
	go func() {
		incomingState, err := state.ParseNetworkState(buf)
		if err != nil {
			panic(err)
		}
		// Get the signatures and rebuild the state
		sigList := incomingState.GetSignatureList()
		for _, sig := range sigList {
			d.peer.GetState().UpdateState(sig)
		}
	}()
}
