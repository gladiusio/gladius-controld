package state

import (
	"encoding/json"
	"errors"
	"reflect"
	"sync"

	"github.com/buger/jsonparser"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
)

// State is a type that represents the network state
type State struct {
	PoolData    *PoolData            `json:"pool_data"`
	NodeDataMap map[string]*NodeData `json:"node_data_map"`
	mux         sync.Mutex
}

// GetJSON gets the JSON representation of the state including signatures
func (s *State) GetJSON() ([]byte, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	return json.Marshal(s)
}

type sigList struct {
	sigs map[string]*signature.SignedMessage
}

func (s *sigList) Add(sig *signature.SignedMessage) {
	if sig != nil {
		s.sigs[string(sig.Hash)] = sig
	}
}

func (s *sigList) GetList() (values []*signature.SignedMessage) {
	for _, v := range s.sigs {
		values = append(values, v)
	}
	return values
}

// GetNodeFields gets the same field from all nodes
func (s *State) GetNodeFields(key string) []interface{} {
	toReturn := make([]interface{}, 0)
	for _, value := range s.NodeDataMap {
		v := reflect.ValueOf(*value)
		toReturn = append(toReturn, v.FieldByName(key).Interface())
	}
	return toReturn
}

func (s *State) GetNodeField(address, key string) interface{} {
	node := s.NodeDataMap[address]
	v := reflect.ValueOf(*node)
	return v.FieldByName(key).Interface()
}

// GetSignatureList returns a list of all of the signed messages used to make
// the current state
func (s *State) GetSignatureList() []*signature.SignedMessage {
	s.mux.Lock()
	defer s.mux.Unlock()
	sigs := &sigList{sigs: make(map[string]*signature.SignedMessage)}

	if s.PoolData != nil {
		// Get all of the pool signatures
		if s.PoolData.FirewallRules != nil {
			for _, field := range s.PoolData.FirewallRules {
				sigs.Add(field.SignedMessage)
			}
		}
		sigs.Add(s.PoolData.RequiredContent.SignedMessage)
	}
	// Get all of the node signatures
	for _, nd := range s.NodeDataMap {
		sigs.Add(nd.LastHeartbeat.SignedMessage)
		sigs.Add(nd.DiskContent.SignedMessage)
		sigs.Add(nd.IPAddress.SignedMessage)
	}

	return sigs.GetList()
}

// UpdateState updates the local state with the signed message information
func (s *State) UpdateState(sm *signature.SignedMessage) (bool, error) {
	if sm.IsInPoolAndVerified() {
		jsonBytes, err := sm.Message.MarshalJSON()
		if err != nil {
			return false, errors.New("malformed state message")
		}

		messageBytes, _, _, err := jsonparser.Get(jsonBytes, "content")
		if err != nil {
			return false, errors.New("can't find content in request")
		}

		timestamp := sm.GetTimestamp()
		updated := false

		handler := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			switch string(key) {
			case "node":
				s.mux.Lock()
				if s.nodeHandler(value, timestamp, sm) {
					updated = true
				}
				s.mux.Unlock()
			case "pool":
				s.mux.Lock()
				if s.poolHandler(value, timestamp, sm) {
					updated = true
				}
				s.mux.Unlock()
			}
			return nil
		}
		jsonparser.ObjectEach(messageBytes, handler)
		return updated, nil
	}
	return false, errors.New("message is not verified")
}

func (s *State) nodeHandler(nodeUpdate []byte, timestamp int64, sm *signature.SignedMessage) bool {
	if s.NodeDataMap == nil {
		s.NodeDataMap = make(map[string]*NodeData)
	}
	if s.NodeDataMap[sm.Address] == nil {
		s.NodeDataMap[sm.Address] = &NodeData{}
	}
	// Keep track of if we update the state or not
	updated := false
	handler := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		switch string(key) {
		case "ip_address":
			// Verify that the timestamp is newer on the incoming signed message
			if s.NodeDataMap[sm.Address].IPAddress.SignedMessage == nil ||
				s.NodeDataMap[sm.Address].IPAddress.SignedMessage.GetTimestamp() < timestamp {
				updated = true

				s.NodeDataMap[sm.Address].IPAddress = SignedField{Data: string(value), SignedMessage: sm}
			}
		case "disk_content":
			// Verify that the timestamp is newer on the incoming signed message
			if s.NodeDataMap[sm.Address].DiskContent.SignedMessage == nil ||
				s.NodeDataMap[sm.Address].DiskContent.SignedMessage.GetTimestamp() < timestamp {
				contentList := make([]string, 0)
				// Get all file names passed in
				jsonparser.ArrayEach(value, func(v []byte, dataType jsonparser.ValueType, offset int, err error) {
					contentList = append(contentList, string(v))
				})

				updated = true

				s.NodeDataMap[sm.Address].DiskContent = SignedList{Data: contentList, SignedMessage: sm}
			}
		case "heartbeat":
			// Verify that the timestamp is newer on the incoming signed message
			if s.NodeDataMap[sm.Address].LastHeartbeat.SignedMessage == nil ||
				s.NodeDataMap[sm.Address].LastHeartbeat.SignedMessage.GetTimestamp() < timestamp {
				updated = true

				s.NodeDataMap[sm.Address].LastHeartbeat = SignedField{Data: string(value), SignedMessage: sm}
			}
		}
		return nil
	}
	jsonparser.ObjectEach(nodeUpdate, handler)
	return updated
}

func (s *State) poolHandler(poolUpdate []byte, timestamp int64, sm *signature.SignedMessage) bool {
	return false
}

// PoolData is a type that stores information about the pool
type PoolData struct {
	FirewallRules   []SignedField `json:"firewall_rules"`
	RequiredContent SignedField   `json:"required_content"`
}

// NodeData is a type that stores infomration about an indiviudal node
type NodeData struct {
	IPAddress     SignedField `json:"ip_address"`
	LastHeartbeat SignedField `json:"last_heartbeat"`
	DiskContent   SignedList  `json:"disk_content"`
}

// SignedField is a type that represents a string field that includes the
// signature that last updated it
type SignedField struct {
	Data          string                   `json:"data"`
	SignedMessage *signature.SignedMessage `json:"signed_message"`
}

// SignedList is a type that represents a list of string fields and includes the
// signature that last updated it
type SignedList struct {
	Data          []string                 `json:"data"`
	SignedMessage *signature.SignedMessage `json:"signed_message"`
}

// ParseNetworkState takes the network state json string in and returns a state
// type if it is valid.
func ParseNetworkState(stateString []byte) (*State, error) {
	s := &State{}
	err := json.Unmarshal(stateString, s)
	return s, err
}
