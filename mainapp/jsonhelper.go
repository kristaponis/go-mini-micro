package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// readJSON
func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have single json value")
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Check for headers
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func errorJSON(w http.ResponseWriter, err error, statusCodes ...int) error {
	status := http.StatusBadRequest

	if len(statusCodes) > 0 {
		status = statusCodes[0]
	}

	jsonRes := jsonResponse{
		Error:   true,
		Message: err.Error(),
	}

	return writeJSON(w, status, jsonRes)
}
