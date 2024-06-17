package main

import (
	"fmt"
	"net/http"
	// "sync"
	"sync/atomic"
)

var uniqueVisits uint64

func main() {
	// mu := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// mu.Lock()
		atomic.AddUint64(&uniqueVisits, 1)
		// mu.Unlock()
		w.Write([]byte(fmt.Sprintf("You're visitor number %d", uniqueVisits)))
	})

	http.ListenAndServe(":3000", nil)
}
