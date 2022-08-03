package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

func main() {
	hotels := []Hotel{
		Hotel{
			Name:    "terry",
			Address: "Gangnam",
			City:    "Seoul",
			Zip:     38342,
			Region:  "Central",
		},

		Hotel{
			Name:    "lu",
			Address: "Songdo",
			City:    "Incheon",
			Zip:     12345,
			Region:  "Northern",
		},

		Hotel{
			Name:    "Sophia",
			Address: "Haewondae",
			City:    "Busan",
			Zip:     99999,
			Region:  "Southern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)

	if err != nil {
		log.Panic(err)
	}

}
