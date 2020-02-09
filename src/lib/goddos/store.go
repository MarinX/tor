package main

import (
	"sync"
	"time"
)

// Store stores data about circuits
type Store struct {
	sync.Mutex
	circuits map[int]*Circuit
}

// Update the circuit
func (s *Store) Update(circ *Circuit) {
	s.Lock()
	defer s.Unlock()
	s.circuits[circ.ID] = circ
}

// GetCircuit returns circuit by ID
func (s *Store) GetCircuit(id int) *Circuit {
	s.Lock()
	defer s.Unlock()
	// if we dont have a circuit in store, create it
	if s.circuits[id] == nil {
		s.circuits[id] = &Circuit{
			ID:      id,
			Counter: NewRateCounter(1 * time.Second),
			Created: time.Now(),
		}
	}
	return s.circuits[id]
}

// Remove circuit from our store
func (s *Store) Remove(id int) {
	s.Lock()
	defer s.Unlock()
	delete(s.circuits, id)
}
