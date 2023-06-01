package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(responseWriter http.ResponseWriter, code int, payload interface{}) {
	json, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failled to marshal JSON response: %v", payload)
		responseWriter.WriteHeader(500)
		return
	}

	responseWriter.Header().Add("Content-Type", "aplication/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(json)
}

func respondWithErro(responseWriter http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Erro interno de servidor...", msg)
	}
	type errResponse struct {
		Erro string `json:"error"`
	}

	respondWithJson(responseWriter, code, errResponse{
		Erro: msg,
	})

}
