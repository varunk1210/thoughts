package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func getThoughts(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"mod":  func(i, j int) int { return i % j },
		"add1": func(i int) int { return i + 1 },
		"sub": func(i, j int) int {
			return i - j
		},
	}
	tmplPath, _ := os.Getwd()
	temp, err := template.New("homepage.tmpl").Funcs(funcMap).ParseFiles(tmplPath + "/Templates/homepage.tmpl")

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
