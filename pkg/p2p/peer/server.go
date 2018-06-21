package peer

import "sync"

// server is a type that manages the Gladius p2p "server"
type server struct {
	running bool
	mux     sync.Mutex
}

// Start starts the p2p server
func (s *server) Start() {

}

// Stop stops the p2p server
func (s *server) Stop() {

}
