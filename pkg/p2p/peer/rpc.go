package peer

import (
	"errors"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
)

type RPCState struct {
	p Peer
}

type RPCLogging struct {
	p Peer
}

func (s *RPCState) Update(arg *signature.SignedMessage, reply *string) error {
	if arg.IsInPoolAndVerified() {
		s.p.UpdateAndPushState(arg)
		*reply = "State Updated"
	} else {
		*reply = "Invalid SignedMessage"
	}
	return nil
}

func (s *RPCState) Get(arg *signature.SignedMessage, reply *string) error {
	if arg.IsVerified() {
		jsonBytes, err := arg.Message.MarshalJSON()
		if err != nil {
			*reply = "Error decoding message"
			return err
		}
		timestamp, err := jsonparser.GetInt(jsonBytes, "content.challenge_time")
		if err != nil {
			*reply = "Can't find field `content.challenge_time`"
			return err
		}
		now := time.Now().Unix()

		if (now - timestamp) < 2 {
			jsonstring, err := s.p.GetState().GetJSON()
			if err != nil {
				*reply = "Error getting full state"
				return err
			}
			*reply = string(jsonstring)
		}
		*reply = "Challenge exprired"
		return errors.New("exprired challenge")
	}
	*reply = "Message not verified"
	return errors.New("message not verified")
}

func (l *RPCLogging) PassMessage(arg *signature.SignedMessage, reply *string) {
	if arg.IsInPoolAndVerified() {
		// TODO: Pass message and don't log it on this machine unless a logging flag is specified in the config
		*reply = "Passed Message"
	}
}
