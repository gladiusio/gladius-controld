package peer

import (
	"errors"

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

func (state *StatePlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *messages.StateMessage:
		sm, err := parseSignedMessage(msg.Message)
		if err == nil {
			state.peerState.UpdateState(sm)
		}
	}

	return nil
}

func (state *StatePlugin) Startup(net *network.Network) {
	// TODO: Create a push/pull sync for missed messages
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
