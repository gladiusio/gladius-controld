package peer

import (
	"encoding/json"

	"github.com/hashicorp/memberlist"
	uuid "github.com/satori/go.uuid"
)

type mergeDelegate struct {
	peer *Peer
}

// NotifyMerge will be called when a join merge event is invoked. It will
// challenge all nodes in the incoming cluster and verify that they are allowed
// into the network by sending them a challenge that they must sign with their
// Ethereum key.
func (md *mergeDelegate) NotifyMerge(peers []*memberlist.Node) error {
	questions := make([]string, len(peers))

	for i, peer := range peers {
		questionString := uuid.NewV4().String()
		questions[i] = questionString

		challenge := &challenge{question: questionString}
		challengeBytes, _ := json.Marshal(challenge)
		md.peer.member.SendReliable(peer, challengeBytes)
	}

	return nil
}
