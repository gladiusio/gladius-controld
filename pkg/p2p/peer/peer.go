package peer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/deckarep/golang-set"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/p2p/message"
	"github.com/gladiusio/gladius-controld/pkg/p2p/peer/messages"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/gladiusio/gladius-controld/pkg/p2p/state"
	"github.com/perlin-network/noise/crypto/ed25519"
	"github.com/perlin-network/noise/network"
	"github.com/perlin-network/noise/network/backoff"
	"github.com/perlin-network/noise/network/discovery"
	"github.com/spf13/viper"
)

// New returns a new peer type
func New(ga *blockchain.GladiusAccountManager) *Peer {
	// TODO: Make this use ethereum keys from the GladiusAccountManager
	keys := ed25519.RandomKeyPair()

	// Build the network
	builder := network.NewBuilder()
	builder.SetKeys(keys)

	// Use KCP instead of TCP, and config bind address and bind port
	builder.SetAddress(network.FormatAddress(
		"kcp",
		viper.GetString("P2P.BindAddress"),
		uint16(viper.GetInt("P2P.BindPort"))))

	// Setup our state and register accepted fields
	s := state.New()
	s.RegisterNodeSingleFields("ip_address", "content_port", "heartbeat")
	s.RegisterNodeListFields("disk_content")

	s.RegisterPoolListFields("required_content")

	// Register peer discovery plugin.
	// TODO: Setup an authorized DHT plugin
	builder.AddPlugin(new(discovery.Plugin))

	// Add the exponential backoff plugin
	builder.AddPlugin(new(backoff.Plugin))

	// Create our state plugin
	statePlugin := new(StatePlugin)
	statePlugin.peerState = s
	builder.AddPlugin(statePlugin)

	net, err := builder.Build()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	go net.Listen()

	peer := &Peer{
		ga:        ga,
		peerState: s,
		net:       net,
		running:   true,
		mux:       sync.Mutex{},
	}
	return peer
}

// Peer is a type that represents a peer in the Gladius p2p network.
type Peer struct {
	ga        *blockchain.GladiusAccountManager
	peerState *state.State
	net       *network.Network
	running   bool
	mux       sync.Mutex
}

// Used to send to a node through an "update"
type challenge struct {
	ChallengeID string
	Question    string
}

// Join will request to join the network from a specific node
func (p *Peer) Join(addressList []string) error {
	for _, addrString := range addressList {
		addr, err := network.ParseAddress(addrString)
		if err != nil {
			return fmt.Errorf("address must look like kcp://host:port, you have %s", addr)
		}
		if addr.Protocol != "kcp" {
			return fmt.Errorf("protocol must be kcp, you have %s", addr.Protocol)
		}
	}
	p.net.Bootstrap(addressList...)
	go func() {
		time.Sleep(1 * time.Second)
		p.net.BroadcastByAddresses(&messages.SyncRequest{}, addressList...)
	}()
	return nil
}

func (p *Peer) UnlockWallet(password string) error {
	_, err := p.ga.UnlockAccount(password)
	return err
}

// SignMessage signs the message with the peer's internal account manager
func (p *Peer) SignMessage(m *message.Message) (*signature.SignedMessage, error) {
	return signature.CreateSignedMessage(m, p.ga)
}

// Stop will stop the peer
func (p *Peer) Stop() {
	p.net.Close()
}

func (p *Peer) SetState(s *state.State) {
	p.mux.Lock()
	p.peerState = s
	p.mux.Unlock()
}

// UpdateAndPushState updates the local state and pushes it to several other peers
func (p *Peer) UpdateAndPushState(sm *signature.SignedMessage) error {
	err := p.GetState().UpdateState(sm)
	if err != nil {
		return err
	}

	signedBytes, err := json.Marshal(sm)
	if err != nil {
		return err
	}

	toSend := &messages.StateMessage{Message: string(signedBytes)}

	p.net.Broadcast(toSend)

	return nil
}

// GetState returns the current local state
func (p *Peer) GetState() *state.State {
	return p.peerState
}

// CompareContent compares the content provided with the content in the state
// and returns a list of the missing files names in the format of:
// website/<"asset" or "route">/filename
func (p *Peer) CompareContent(contentList []string) []interface{} {
	// Convert to an interface array
	cl := make([]interface{}, len(contentList))
	for i, v := range contentList {
		cl[i] = v
	}
	contentWeHaveSet := mapset.NewSetFromSlice(cl)

	contentField := p.GetState().GetPoolField("required_content")
	if contentField == nil {
		return make([]interface{}, 0)
	}
	contentFromPool := contentField.(*state.SignedList).Data

	// Convert to an interface array
	s := make([]interface{}, len(contentFromPool))
	for i, v := range contentFromPool {
		s[i] = v
	}

	// Create a set
	contentWeNeed := mapset.NewSetFromSlice(s)

	// Return the difference of the two
	return contentWeNeed.Difference(contentWeHaveSet).ToSlice()
}

// GetContentLinks returns a map mapping a file name to all of the URLS it can
// be found on from the network
func (p *Peer) GetContentLinks(contentList []string) map[string][]string {
	allContent := p.GetState().GetNodeFieldsMap("disk_content")
	toReturn := make(map[string][]string)
	for nodeAddress, diskContent := range allContent {
		ourContent := diskContent.(*state.SignedList).Data
		// Convert to an interface array
		s := make([]interface{}, len(ourContent))
		for i, v := range ourContent {
			s[i] = v
		}
		ourContentSet := mapset.NewSetFromSlice(s)
		// Check to see if the current node we're iterating over has any of the
		// content we want
		for _, contentWanted := range contentList {
			if ourContentSet.Contains(contentWanted) {
				if toReturn[contentWanted] == nil {
					toReturn[contentWanted] = make([]string, 0)
				}
				// Add the URL to the map
				link := p.createContentLink(nodeAddress, contentWanted)
				if link != "" {
					toReturn[contentWanted] = append(toReturn[contentWanted], link)
				}
			}
		}
	}
	return toReturn
}

// Builds a URL to a node
func (p *Peer) createContentLink(nodeAddress, contentFileName string) string {
	peerState := p.GetState()
	if p.GetState() == nil {
		return ""
	}

	nodeIPField := peerState.GetNodeField(nodeAddress, "ip_address")
	nodePortField := peerState.GetNodeField(nodeAddress, "content_port")
	if nodeIPField == nil || nodePortField == nil {
		return ""
	}

	nodeIP := nodeIPField.(*state.SignedField).Data
	nodePort := nodePortField.(*state.SignedField).Data
	if nodeIP == nil || nodePort == nil {
		return ""
	}

	contentData := strings.Split(contentFileName, "/")
	u := url.URL{}

	u.Host = nodeIP.(string) + ":" + nodePort.(string)
	u.Path = "/content"
	u.Scheme = "http"

	if len(contentData) == 2 {
		q := u.Query()
		q.Add("website", contentData[0]) // website name
		q.Add("asset", contentData[1])   // "asset" to name of file
		u.RawQuery = q.Encode()
		return u.String()
	}
	return ""
}

// GetContentLocations returns a map mapping a file name to all the content nodes it can
// be found on in the network
func (p *Peer) GetContentLocations(contentList []string) map[string][]interface{} {
	nodeMap := p.GetState().GetNodeMultipleFieldsMap("ip_address", "content_port", "disk_content")
	toReturn := make(map[string][]interface{}) // map file name to array of nodes
	for nodeAddress, nodeInfo := range nodeMap {

		ourContent := nodeInfo["disk_content"].(*state.SignedList).Data
		// Convert to an interface array
		s := make([]interface{}, len(ourContent))
		for i, v := range ourContent {
			s[i] = v
		}
		ourContentSet := mapset.NewSetFromSlice(s)
		// Check to see if the current node we're iterating over has any of the
		// content we want
		for _, contentWanted := range contentList {
			if ourContentSet.Contains(contentWanted) {
				if toReturn[contentWanted] == nil {
					toReturn[contentWanted] = make([]interface{}, 0)
				}
				// Add the node info to the map
				toReturn[contentWanted] = append(toReturn[contentWanted],
					nodeMap[nodeAddress])
			}
		}
	}
	return toReturn
}
