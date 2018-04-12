package models

import (
	"fmt"
)

// Market struct to map to the Market Smart Contract Structure
type Market struct {
	address string
	owner   string
}

func (m *Market) getAddress() string {
	fmt.Println("Market addres: ", m.address)
	return m.address
}

func (m *Market) pools() []string {
	var pools []string
	return pools
}

func (m *Market) ownedPools() []string {
	var pools []string
	return pools
}

func (m *Market) createPool() bool {
	return false
}

func (m *Market) allocateFunds(address string) bool {
	return false
}
