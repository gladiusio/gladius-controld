package peer

import (
	"log"
	"net"
	"net/http"
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
	rpcState := &RPCState{p: s.peer}
	rpc.RegisterName("State", rpcState)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":4351")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(l, nil)
}
