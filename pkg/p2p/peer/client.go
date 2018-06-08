package peer

import "sync"

// client is a type that manages the Gladius peer "client"
type client struct {
	running bool
	mux     sync.Mutex
}

// Start starts the p2p client
func (s *client) Start() {

}

// Stop stops the p2p client
func (s *client) Stop() {

}
