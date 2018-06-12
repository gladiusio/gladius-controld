package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/buger/jsonparser"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
)

// State is a type that represents the network state
type State struct {
	PoolData     PoolData   `json:"pool_data"`
	NodeDataList []NodeData `json:"node_data_list"`
}

// GetJSON gets the JSON representation of the state including signatures
func (s State) GetJSON() ([]byte, error) {
	return json.Marshal(s)
}

// UpdateState updates the local state with the signed message information
func (s *State) UpdateState(sm *signature.SignedMessage) {
	if sm.IsVerified() {
		jsonBytes, err := sm.Message.MarshalJSON()
		if err != nil {
			log.Println(errors.New("Malformed state JSON"))
			return
		}

		messageBytes, _, _, err := jsonparser.Get(jsonBytes, "content")
		if err != nil {
			log.Println("Couldn't process state update")
			return
		}

		var handler func([]byte, []byte, jsonparser.ValueType, int) error
		handler = func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			switch string(key) {
			case "node":
				nodeHandler(value)
			}
			return nil
		}
		jsonparser.ObjectEach(messageBytes, handler)

		jsonparser.Get(jsonBytes)
	}
}

func nodeHandler(nodeUpdate []byte) {
	fmt.Println(string(nodeUpdate))
}

// PoolData is a type that stores information about the pool
type PoolData struct {
	FirewallRules   []SignedField `json:"firewall_rules"`
	RequiredContent SignedField   `json:"required_content"`
}

// NodeData is a type that stores infomration about an indiviudal node
type NodeData struct {
	IPAddress     SignedField   `json:"ip_address"`
	LastHeartbeat SignedField   `json:"last_heartbeat"`
	DiskContent   []SignedField `json:"disk_content"`
}

// SignedField is a type that represents a string field that includes the
// signature that last updated it
type SignedField struct {
	Data          string                  `json:"data"`
	SignedMessage signature.SignedMessage `json:"signed_message"`
}

// ParseNetworkState takes the network state json string in and returns a state
// type if it is valid.
func ParseNetworkState(stateString []byte) (*State, error) {
	return &State{}, nil
}
