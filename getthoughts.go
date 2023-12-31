package main

import (
	"html/template"
	"log"
	"net/http"
)

func getThoughts(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"mod":  func(i, j int) int { return i % j },
		"add1": func(i int) int { return i + 1 },
		"sub": func(i, j int) int {
			return i - j
		},
	}

	temp, err := template.New("homepage.tmpl").Funcs(funcMap).ParseFiles("templates/homepage.tmpl")

	if err != nil {
		log.Println("Error parsing template: ", err)
		return
	}

	// getting all the documents from the database
	getAllDocuments := GetAllDocuments()

	err = temp.Execute(w, getAllDocuments)
	if err != nil {
		log.Println("Error executing template: ", err)
	}

}
