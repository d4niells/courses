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
	tmp := template.New("CourseTemplate")
	tmp, err := tmp.Parse("Course: {{.Name}} - Duration: {{.Duration}} \n")
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(os.Stdout, course)
}
