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
	course := Course{"Go expert", 44}
	t := template.Must(template.New("CursoTemplate").Parse("Course: {{.Name}} - Duration: {{.Duration}}\n"))

	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}

}
