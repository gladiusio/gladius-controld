package server

import "sync"

// Server is a type that manages the Gladius p2p interface
type Server struct {
	running bool
	mux     sync.Mutex
}

// Start starts the p2p server
func (s *Server) Start() {

}

// Stop stops the p2p server
func (s *Server) Stop() {

}
