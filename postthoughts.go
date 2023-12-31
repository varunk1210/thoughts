package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func PostThoughts(w http.ResponseWriter, r *http.Request) {
	// Post thoughts code goes here
	log.Println("HTMX request received.")
	log.Println(r.Header.Get("HX-Request"))

	// Data to the database
	thought := r.FormValue("thought")
	currentTime := time.Now().UTC().Format("2006-01-02 15:04:05")
	// All the values are set

	log.Println("Thought received: ", thought)
	data := bson.M{
		"thought": thought, "time": currentTime,
	}

	// Inserting the data into the database
	err := InsertDocument(data)
	if err != nil {
		log.Println("Error inserting document: ", err)
	}

	uploadData := ThoughtsData{
		Thought: thought, Time: currentTime}

	log.Println(uploadData.Time)

	funcMap := template.FuncMap{
		"mod":  func(i, j int) int { return i % j },
		"add1": func(i int) int { return i + 1 },
		"add": func(a, b int) int {
			return a + b
		},
	}

	// getting all the documents from the database
	getAllDocuments := GetAllDocuments()
	length := len(getAllDocuments)

	log.Printf("%v\n\n\n", length)

	//htmlstring
	htmlStr := fmt.Sprintf("<div class=\"col\"> <div class=\"card\"> <div class=\"card-body\"> <h4 class=\"card-title\">Thought - %v </h4> <h6 class=\"text-muted card-subtitle mb-2\">%v</h6> <p class=\"card-text\">%v</p> </div> </div> </div>", length, uploadData.Time, uploadData.Thought)
	// Parsing the template
	template, _ := template.New("t").Funcs(funcMap).Parse(htmlStr)

	// Executing the template
	err = template.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template: ", err)
	}

	log.Println("Thoughts posted.")
}
