package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) GetName() string {
	return p.Name
}

func (p person) GetAge() int {
	return p.Age
}

func (p person) SumAge(extra int) int {
	return p.Age + extra
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}
func main() {
	p := person{
		Name: "terry",
		Age:  25,
	}
	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Panic(err)
	}
}
