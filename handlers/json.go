package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Println("Failed to marshal json", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("Server Error:", msg)
	}
	type ErrorMsg struct {
		Message string `json:"error"`
	}
	respondWithJson(w, statusCode, ErrorMsg{Message: msg})
}
