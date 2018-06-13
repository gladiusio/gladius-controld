package peer

import (
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
)

type RPCState struct {
	p Peer
}

func (s *RPCState) Update(arg *signature.SignedMessage, reply *string) error {
	if arg.IsVerified() {
		s.p.UpdateAndPushState(arg)
		*reply = "State Updated"
	} else {
		*reply = "Invalid State"
	}
	return nil
}

func (s *RPCState) Get(args int, reply *string) error {
	jsonstring, err := s.p.GetState().GetJSON()
	*reply = string(jsonstring)
	return err
}
