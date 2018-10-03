package state

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/buger/jsonparser"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
)

// State is a type that represents the network state
type State struct {
	// poolDataFields and nodeDataFields keep track of what fields are valid for
	// the protocol, 0 is SignedField, 1 is SignedList
	poolDataFields map[string]int
	nodeDataFields map[string]int

	// Keeps track of the actual data
	PoolData    PoolData            `json:"pool_data"`
	NodeDataMap map[string]NodeData `json:"node_data_map"`

	mux sync.Mutex
}

// New returns a pointer to a State object
func New() *State {
	s := &State{}
	s.poolDataFields = make(map[string]int)
	s.nodeDataFields = make(map[string]int)
	return s
}

// RegisterPoolListFields registers the fields as understood types to be recorded in
// the state of the pool
func (s *State) RegisterPoolListFields(fields ...string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, field := range fields {
		s.poolDataFields[field] = 1
	}
}

// RegisterPoolSingleFields registers the fields as understood types to be recorded in
// the state of the pool
func (s *State) RegisterPoolSingleFields(fields ...string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, field := range fields {
		s.poolDataFields[field] = 0
	}
}

// RegisterNodeListFields registers the fields as understood types to be recorded in
// the state of a node
func (s *State) RegisterNodeListFields(fields ...string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, field := range fields {
		s.nodeDataFields[field] = 1
	}
}

// RegisterNodeSingleFields registers the fields as understood types to be recorded in
// the state of a node
func (s *State) RegisterNodeSingleFields(fields ...string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, field := range fields {
		s.nodeDataFields[field] = 0
	}
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

// GetPoolField gets the field by name from the pool
func (s *State) GetPoolField(key string) interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.PoolData != nil {
		return s.PoolData[key]
	}
	return nil
}

// GetNodeFields gets the same field from all nodes
func (s *State) GetNodeFields(key string) []interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()

	toReturn := make([]interface{}, 0)
	for _, node := range s.NodeDataMap {
		toReturn = append(toReturn, node[key])
	}
	return toReturn
}

// GetNodeFieldsMap gets a map of node address to the field referenced by key
func (s *State) GetNodeFieldsMap(key string) map[string]interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()

	toReturn := make(map[string]interface{})

	// Go through every node and get that specific field
	for node, data := range s.NodeDataMap {
		if data[key] != nil {
			toReturn[node] = data[key]
		}
	}
	return toReturn
}

// GetNodeMultipleFieldsMap gets a map of node address to the fields referenced by keys
func (s *State) GetNodeMultipleFieldsMap(keys ...string) map[string]map[string]interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()

	nodeMap := make(map[string]map[string]interface{})

	// Go through every node and get the fields requested by keys
	for node, data := range s.NodeDataMap {
		newNode := make(map[string]interface{})
		for _, key := range keys {
			if data[key] != nil {
				newNode[key] = data[key]
			}
		}
		nodeMap[node] = newNode
	}
	return nodeMap
}

func (s *State) GetNodeField(address, key string) interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.NodeDataMap[address][key]
}

// GetSignatureList returns a list of all of the signed messages used to make
// the current state
func (s *State) GetSignatureList() []*signature.SignedMessage {
	s.mux.Lock()
	defer s.mux.Unlock()
	sigs := &sigList{sigs: make(map[string]*signature.SignedMessage)}

	if s.PoolData != nil {
		for _, field := range s.PoolData {
			switch typedField := field.(type) {
			case *SignedList:
				sigs.Add(typedField.SignedMessage)
			case *SignedField:
				sigs.Add(typedField.SignedMessage)
			}
		}
	}
	// Get all of the node signatures
	for _, nd := range s.NodeDataMap {
		for _, field := range nd {
			switch typedField := field.(type) {
			case *SignedList:
				sigs.Add(typedField.SignedMessage)
			case *SignedField:
				sigs.Add(typedField.SignedMessage)
			}
		}
	}

	return sigs.GetList()
}

// UpdateState updates the local state with the signed message information
func (s *State) UpdateState(sm *signature.SignedMessage) error {
	if sm.IsInPoolAndVerified() {
		jsonBytes, err := sm.Message.MarshalJSON()
		if err != nil {
			return errors.New("malformed state message")
		}

		messageBytes, _, _, err := jsonparser.Get(jsonBytes, "content")
		if err != nil {
			return errors.New("can't find content in request")
		}

		timestamp := sm.GetTimestamp()

		handler := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			var err error
			suc := false
			switch string(key) {
			case "node":
				s.mux.Lock()
				suc, err = s.nodeHandler(value, timestamp, sm)
				s.mux.Unlock()
			case "pool":
				s.mux.Lock()
				suc, err = s.poolHandler(value, timestamp, sm)
				s.mux.Unlock()
			}
			if !suc && err == nil {
				return errors.New("Nothing was updated")
			}
			return err
		}
		return jsonparser.ObjectEach(messageBytes, handler)
	}
	return errors.New("message is not verified")
}

func (s *State) isUnderstoodNodeField(key string) bool {
	_, ok := s.nodeDataFields[key]
	return ok
}

func (s *State) fieldType(key string) int {
	if _, ok := s.nodeDataFields[key]; ok {
		return s.nodeDataFields[key]
	}
	return s.poolDataFields[key]
}

func (s *State) isUnderstoodPoolField(key string) bool {
	_, ok := s.poolDataFields[key]
	return ok
}

func (s *State) nodeHandler(nodeUpdate []byte, timestamp int64, sm *signature.SignedMessage) (bool, error) {
	if s.NodeDataMap == nil {
		s.NodeDataMap = make(map[string]NodeData)
	}
	if s.NodeDataMap[sm.Address] == nil {
		s.NodeDataMap[sm.Address] = NodeData{}
	}
	// Keep track of if we update the state or not
	updated := false
	handler := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		keyString := string(key)
		// If it's a different protocol, or not an understood field, don't add it to
		// our state
		if s.isUnderstoodNodeField(keyString) {
			// Check if the our node has never been updated, or the incomming message
			// is newer than the one we have
			if s.fieldType(keyString) == 0 {
				if s.NodeDataMap[sm.Address][keyString] == nil ||
					s.NodeDataMap[sm.Address][keyString].(*SignedField).SignedMessage.GetTimestamp() < timestamp {

					// Actually update the field
					s.NodeDataMap[sm.Address][keyString] = &SignedField{Data: string(value), SignedMessage: sm}
					updated = true
					return nil
				}
			} else {
				if s.NodeDataMap[sm.Address][keyString] == nil ||
					s.NodeDataMap[sm.Address][keyString].(*SignedList).SignedMessage.GetTimestamp() < timestamp {

					// Create a string list
					contentList := make([]string, 0)
					// Get all file names passed in
					jsonparser.ArrayEach(value, func(v []byte, dataType jsonparser.ValueType, offset int, err error) {
						contentList = append(contentList, string(v))
					})

					// Actually update the field
					s.NodeDataMap[sm.Address][keyString] = &SignedList{Data: contentList, SignedMessage: sm}
					updated = true
					return nil
				}
			}
			return errors.New("Message was older than the current version")
		}
		return errors.New("Unsupported field in update message")
	}
	err := jsonparser.ObjectEach(nodeUpdate, handler)
	return updated, err
}

func (s *State) poolHandler(poolUpdate []byte, timestamp int64, sm *signature.SignedMessage) (bool, error) {
	if s.PoolData == nil {
		s.PoolData = PoolData{}
	}

	// Don't update the state
	if !sm.IsPoolManagerAndVerified() {
		return false, errors.New("Message is not verified or not from pool manager")
	}

	// Keep track of if we update the state or not
	updated := false
	handler := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		keyString := string(key)
		// If it's a different protocol, or not an understood field, don't add it to
		// our state
		if s.isUnderstoodPoolField(keyString) {
			// Check if the our node has never been updated, or the incomming message
			// is newer than the one we have
			if s.fieldType(keyString) == 0 {
				if s.PoolData[keyString] == nil ||
					s.PoolData[keyString].(*SignedField).SignedMessage.GetTimestamp() < timestamp {

					// Actually update the field
					s.PoolData[keyString] = &SignedField{Data: string(value), SignedMessage: sm}
					updated = true
					return nil
				}
			} else {
				if s.PoolData[keyString] == nil ||
					s.PoolData[keyString].(*SignedList).SignedMessage.GetTimestamp() < timestamp {
					// Create a string list
					contentList := make([]string, 0)
					// Get all file names passed in
					jsonparser.ArrayEach(value, func(v []byte, dataType jsonparser.ValueType, offset int, err error) {
						contentList = append(contentList, string(v))
					})

					// Actually update the field
					s.PoolData[keyString] = &SignedList{Data: contentList, SignedMessage: sm}
					updated = true
					return nil
				}
			}
			return errors.New("Message was older than the current version")
		}
		return errors.New("Unsupported field in update message")
	}
	err := jsonparser.ObjectEach(poolUpdate, handler)
	return updated, err
}

// PoolData is a type that stores information about the pool
type PoolData map[string]interface{}

// NodeData is a type that stores infomration about an indiviudal node
type NodeData map[string]interface{}

// SignedField is a type that represents a string field that includes the
// signature that last updated it
type SignedField struct {
	Data          interface{}              `json:"data"`
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
