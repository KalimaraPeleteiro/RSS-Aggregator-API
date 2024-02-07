package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

// Criando Usuários
func (apiConfiguration *apiConfig) handlerCreateUser(response http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name     string `json:"nome"`
		Password string `json:"senha"`
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
		Password:  params.Password,
	})

	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Erro ao criar usuário: %s", err))
		return
	}

	JSONResponse(response, 201, SQLCUserToUser(user))

}

// Buscando Usuário
func (apiConfiguration *apiConfig) handlerGetUserByAPIKey(response http.ResponseWriter, request *http.Request, user database.User) {
	JSONResponse(response, 200, SQLCUserToUser(user))
}

// Criando Usuários
func (apiConfiguration *apiConfig) handlerLoginAuthentication(response http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name     string `json:"nome"`
		Password string `json:"senha"`
	}

	decoder := json.NewDecoder(request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Erro ao decodificar JSON: %s", err))
		return
	}

	user, err := apiConfiguration.database.LoginAuthentication(request.Context(), database.LoginAuthenticationParams{
		Name:     params.Name,
		Password: params.Password,
	})

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			errorJSON(response, 404, fmt.Sprintf("Credenciais inválidas."))
			return
		}
		errorJSON(response, 400, fmt.Sprintf("Erro ao autenticar usuário: %s", err))
		return
	}

	JSONResponse(response, 200, user)

}
