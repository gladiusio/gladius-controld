package peer

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/buger/jsonparser"
	"github.com/perlin-network/noise/network"

	"github.com/gladiusio/gladius-controld/pkg/p2p/peer/messages"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/gladiusio/gladius-controld/pkg/p2p/state"
)

type StatePlugin struct {
	*network.Plugin
	peerState *state.State
}

// Receive is called every time a new message is received
func (state *StatePlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *messages.StateMessage:
		sm, err := parseSignedMessage(msg.Message)
		if err == nil {
			go state.peerState.UpdateState(sm)
		}
	case *messages.SyncRequest:
		smList := state.peerState.GetSignatureList()
		smStringList := make([]string, 0)
		for _, sm := range smList {
			smString, _ := json.Marshal(sm)
			smStringList = append(smStringList, string(smString))
		}
		ctx.Reply(&messages.SyncResponse{SignedMessage: smStringList})
	case *messages.SyncResponse:
		smStringList := msg.SignedMessage
		for _, smString := range smStringList {
			sm, err := parseSignedMessage(smString)
			if err != nil {
				return errors.New("Invalid signed message sent")
			}
			go state.peerState.UpdateState(sm)
		}
	}

	return nil
}

// Startup is called once the network is bootstrapped. Every 60
// seconds we ask a random peer for it's state. This is an anti
// entropy method that might not be entirely needed.
func (state *StatePlugin) Startup(net *network.Network) {
	go func() {
		time.Sleep(60 * time.Second)
		net.BroadcastRandomly(&messages.SyncRequest{}, 1)
	}()
}

func parseSignedMessage(sm string) (*signature.SignedMessage, error) {
	smBytes := []byte(sm)

	messageBytes, _, _, err := jsonparser.Get(smBytes, "message")
	if err != nil {
		return nil, errors.New("Can't find `message` in body")
	}

	hash, err := jsonparser.GetString(smBytes, "hash")
	if err != nil {
		return nil, errors.New("Can't find `hash` in body")
	}

	signatureString, err := jsonparser.GetString(smBytes, "signature")
	if err != nil {
		return nil, errors.New("Could not find `signature` in body")

	}

	address, err := jsonparser.GetString(smBytes, "address")
	if err != nil {
		return nil, errors.New("Could not find `address` in body")
	}

	parsed, err := signature.ParseSignedMessage(string(messageBytes), hash, signatureString, address)
	if err != nil {
		return nil, errors.New("Couldn't parse body")

	}

	return parsed, nil
}
