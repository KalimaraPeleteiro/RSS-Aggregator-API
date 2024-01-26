package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Arquivo responsável pela definição de funções que geram as respostas JSON para a API.

func errorJSON(response http.ResponseWriter, statusCode int, message string) {

	// Erros 500+ são problemas do nosso lado, então é melhor manter noção disso.
	if statusCode > 499 {
		log.Printf("Requisição levou a erro 5XX: %v", message)
	}

	type errorMessage struct {
		Error string `json:"Erro"`
	}

	JSONResponse(response, statusCode, errorMessage{
		Error: message,
	})
}

func JSONResponse(response http.ResponseWriter, statusCode int, answer interface{}) {
	data, err := json.Marshal(answer)

	if err != nil {
		log.Printf("O JSON não foi capaz de transformar o conteúdo: %v", answer)
		response.WriteHeader(500)
		return
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(data)
	log.Printf("Respondendo Requisição. Status: %v", statusCode)
}
