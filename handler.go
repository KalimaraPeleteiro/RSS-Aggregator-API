package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

// Arquivo responsável por lidar com as requisições

// Rotas padrão para verificação
func handlerDefaultError(response http.ResponseWriter, request *http.Request) {
	errorJSON(response, 400, "Esta é uma mensagem padrão de erro.")
}

func handlerServerReadiness(response http.ResponseWriter, request *http.Request) {
	JSONResponse(response, 200, struct{}{})
}

// Criando Usuários
func (apiConfiguration *apiConfig) handlerCreateUser(response http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name string `json:"nome"`
	}

	decoder := json.NewDecoder(request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Erro ao decodificar JSON: %s", err))
		return
	}

	user, err := apiConfiguration.database.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Erro ao criar usuário: %s", err))
		return
	}

	JSONResponse(response, 200, SQLCUserToUser(user))

}
