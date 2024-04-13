package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Duration int
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := t.Execute(w, []Course{
			{"Go", 200},
			{"Java", 100},
			{"Python", 400},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", mux)
}
