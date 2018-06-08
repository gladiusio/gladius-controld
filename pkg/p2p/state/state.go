package state

import (
	"encoding/json"

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
func ParseNetworkState(stateString []byte) {

}
