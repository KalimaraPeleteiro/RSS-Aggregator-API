package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extraindo chave dos headers.

// Método correto:
// Authorization: ApiKey chave

// Exemplo:
// ApiKey fd3de1d9e3a9ad88138aba3b73f879b0f3a208240e05140cb125761784358a87
func GetAPIKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")

	if value == "" {
		return "", errors.New("nenhuma chave foi encontrada na requisição.")
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		return "", errors.New("o Header não foi organizado corretamente.")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("primeira parte do Header está errada.")
	}

	return values[1], nil
}
