package peer

import (
	"sync"
)

// client is a type that manages the Gladius peer "client"
type client struct {
	running bool
	mux     sync.Mutex
}
