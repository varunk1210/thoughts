package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	// File server directory
	fs := http.FileServer(http.Dir("./Templates/"))
	router.Handle("/*", fs)

	// Database connection
	client, err := connection()
	if err != nil {
		log.Println("Error connecting to MongoDB: ", err)
	} else {
		log.Println("Connected to MongoDB.")
		defer disconnect(client)
	}

	// Routes
	router.Get("/", getThoughts)
	router.Post("/post-thoughts/", PostThoughts)

	// Port related settings
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8090"
	}

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(port, router))
}
