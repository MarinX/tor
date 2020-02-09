package main


import "C"

import (
	"fmt"
	"sync"
	"time"

)

// feel free to modify variable here to suits your needs
const maxRequestsPerSecond = 3
const banTime = 30 * time.Minute

// global store variable
var store *Store
// do only once
var initWorker sync.Once


//export GoCircuitHandler
func GoCircuitHandler(cid C.int) *C.char {
	initWorker.Do(func(){
		store = &Store{
			circuits: make(map[int]*Circuit),
		}
	})
	
	id := int(cid)

	circ := store.GetCircuit(id)
	circ.AddRequest()

	if circ.IsBanned {
		
		// check if ban expired
		if circ.Created.After(circ.Created.Add(banTime)) {
			circ.Clear()
			store.Update(circ)
			return nil
		}
		return C.CString(fmt.Sprintf("GoDDoS %d already banned. Rate %d req/sec", id, circ.Counter.Rate()))
	}

	if circ.IsMaxOut() {
		circ.IsBanned = true
		store.Update(circ)
		return C.CString(fmt.Sprintf("GoDDoS %d reached max, ban. Rate %d req/sec", id, circ.Counter.Rate()))
	}

	return nil
}

func main() {}
