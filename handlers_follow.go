package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Usuário passa a seguir novo feed
func (apiConfiguration *apiConfig) handlerFollowNewFeed(response http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Erro ao decodificar JSON: %s", err))
		return
	}

	newFollowing, err := apiConfiguration.database.FollowNewFeed(request.Context(), database.FollowNewFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Erro ao seguir nova feed: %s", err))
		return
	}

	JSONResponse(response, 201, SQLCFollowingFeedToFollowingFeed(newFollowing))

}

// Retornando todos os Feeds que um Usuário Segue
func (apiConfiguration apiConfig) handlerGetAllUserFollowingFeeds(response http.ResponseWriter, request *http.Request, user database.User) {
	feeds, err := apiConfiguration.database.ReturnUserFollowingFeeds(request.Context(), user.ID)
	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Não consegui buscar os feeds que você segue: %v", err))
		return
	}

	JSONResponse(response, 200, SQLCFollowingFeedsToFollowingFeeds(feeds))
}

// Retornando todos os Feeds que um Usuário Segue
func (apiConfiguration apiConfig) handlerUnfollowFeed(response http.ResponseWriter, request *http.Request, user database.User) {
	id := chi.URLParam(request, "feed_id")
	feed_id, err := uuid.Parse(id)

	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Houve uma erro na chave de sua feed: %v", err))
		return
	}

	err = apiConfiguration.database.UnfollowFeed(request.Context(), database.UnfollowFeedParams{
		ID:     feed_id,
		UserID: user.ID,
	})

	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Não consegui desfazer o seguimento: %v", err))
		return
	}

	JSONResponse(response, 200, struct{}{})
}
