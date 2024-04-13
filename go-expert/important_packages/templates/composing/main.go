package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Course struct {
	Name     string
	Duration int
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{"header.html", "content.html", "courses.html"}

		funcMap := template.FuncMap{
			"ToUpper": strings.ToUpper,
			"ToTime":  func(n int) string { return fmt.Sprintf("%v hours of content", n) },
		}

		t := template.Must(template.New("content.html").Funcs(funcMap).ParseFiles(templates...))

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
