package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./public")))
	// mux.Handle("/posts", http.FileServer(http.Dir("./public/posts.html")))
	http.ListenAndServe(":8080", mux)
}
