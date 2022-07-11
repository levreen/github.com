package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

/* func main() {
	templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(templateText)
	check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)
	err = tmpl.Execute(os.Stdout, 42)
	check(err)
	err = tmpl.Execute(os.Stdout, true)
	check(err)
}
*/

/*func main() {
	executeTemplate("Dot is: {{if .}}Dot is true!{{end}} finish\n", true)
	executeTemplate("Dot is: {{if .}}Dot is true!{{end}} finish\n", false)
}
*/

/*
func main() {
	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop:{{.}}\n"
	executeTemplate(templateText, []string{"do", "re", "mi"})
}
*/

/*
func main() {
	templateText := "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 0.99, 27})
}
*/

// type Part struct {
// 	Name  string
// 	Count int
// }

// func main() {
// 	templateText := "Name: {{.Name}}\nCount: {{.Count}}\n"
// 	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
// 	executeTemplate(templateText, Part{Name: "Cables", Count: 2})
// }

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func main() {
	templateText := "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Bombay Shit", Rate: 4.99, Active: true}
	executeTemplate(templateText, subscriber)
	subscriber = Subscriber{Name: "Joy Carr", Rate: 5.99, Active: false}
	executeTemplate(templateText, subscriber)
}
