package peer

import "github.com/hashicorp/memberlist"

type mergeDelegate struct {
	peer *Peer
}

func (md *mergeDelegate) NotifyMerge(peers []*memberlist.Node) error {
	return nil
}
