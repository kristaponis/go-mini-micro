package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Connect to database.
	_, err := openDB(os.Getenv("DSN"))
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	// Start the server with custom config.
	srv := &http.Server{
		Addr:    ":80",
		Handler: routes(),
	}

	log.Println("Starting auth service at port 80..")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
