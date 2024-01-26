package main

import "net/http"

// Arquivo responsável por lidar com as requisições

// Rotas padrão para verificação
func handlerDefaultError(response http.ResponseWriter, request *http.Request) {
	errorJSON(response, 400, "Esta é uma mensagem padrão de erro.")
}

func handlerServerReadiness(response http.ResponseWriter, request *http.Request) {
	JSONResponse(response, 200, struct{}{})
}
