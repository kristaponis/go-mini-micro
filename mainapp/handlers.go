package main

import (
	"log"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "hello from main app",
	}

	err := writeJSON(w, http.StatusOK, payload)
	if err != nil {
		log.Println("Error writing JSON:", err)
	}
}
