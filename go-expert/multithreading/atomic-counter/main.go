package main

import (
	"fmt"
	"net/http"
)

var uniqueVisits uint64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uniqueVisits++
		w.Write([]byte(fmt.Sprintf("You're visitor number %d", uniqueVisits)))
	})

	http.ListenAndServe(":3000", nil)
}
