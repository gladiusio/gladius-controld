package peer

import (
	"log"
	"net"
	"net/rpc"
	"sync"
)

// server is a type that manages the Gladius p2p "server"
type server struct {
	running bool
	peer    *Peer
	mux     sync.Mutex
}

func newServer(p *Peer) *server {
	return &server{running: false, peer: p}
}

// Start starts the p2p server
func (s *server) Start() {
	rpcServer := rpc.NewServer()
	rpcState := &RPCState{p: s.peer}
	rpcServer.RegisterName("State", rpcState)
	// Listen for incoming tcp packets on specified port.
	l, e := net.Listen("tcp", ":4351")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// This statement links rpc server to the socket, and allows rpc server to accept
	// rpc request coming from that socket.
	go rpcServer.Accept(l)
}
