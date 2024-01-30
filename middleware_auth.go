package main

import (
	"fmt"
	"net/http"

	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/auth"
	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfiguration apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		key, err := auth.GetAPIKey(request.Header)

		if err != nil {
			errorJSON(response, 403, fmt.Sprintf("Erro de autenticação: %v", err))
			return
		}

		user, err := apiConfiguration.database.GetUseByAPIKey(request.Context(), key)

		if err != nil {
			errorJSON(response, 400, fmt.Sprintf("Não consegui encontrar usuários com essa chave. Talvez você digitou errado?"))
			return
		}

		handler(response, request, user)
	}
}
