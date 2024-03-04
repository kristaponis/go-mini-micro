package main

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func handleMainApp(w http.ResponseWriter, r *http.Request) {
	payload := response{
		Error:   false,
		Message: "hello from main app",
	}

	out, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
