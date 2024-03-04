package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":8000",
		Handler: routes(),
	}

	fmt.Println("Starting main app service on port 8000")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
