package main

import (
	"time"
)

// Circuit holds the data about tor circuit
type Circuit struct {
	ID       int
	Counter  *RateCounter
	IsBanned bool
	Created  time.Time
}

// AddRequest increments time that particular circuit opens the connection
func (c *Circuit) AddRequest() {
	c.Counter.Incr(1)
}

// IsMaxOut returns true if too many request are going through particular circuit per second
func (c *Circuit) IsMaxOut() bool {
	return c.Counter.Rate() > maxRequestsPerSecond
}

// Clear clears the circuit data - this is usefull for unban
func (c *Circuit) Clear() {
	c.Created = time.Now()
	c.IsBanned = false
	c.Counter = NewRateCounter(1 * time.Second)
}
