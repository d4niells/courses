package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Duration int
}

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, []Course{
		{"Go", 200},
		{"Java", 100},
		{"Python", 400},
	})
	if err != nil {
		panic(err)
	}
}
