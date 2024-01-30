package main

import (
	"time"

	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

// Simplesmente para deixar o JSON de resposta traduzido.

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"criado_em"`
	UpdatedAt time.Time `json:"atualizado_em"`
	Name      string    `json:"nome"`
	ApiKey    string    `json:"chave_API"`
}

func SQLCUserToUser(sqlcUser database.User) User {
	return User{
		ID:        sqlcUser.ID,
		CreatedAt: sqlcUser.CreatedAt,
		UpdatedAt: sqlcUser.UpdatedAt,
		Name:      sqlcUser.Name,
		ApiKey:    sqlcUser.ApiKey,
	}
}
