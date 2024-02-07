package main

import (
	"fmt"
	"net/http"

	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/database"
)

// Retornando os Posts que dos Feeds que o User Segue
func (apiConfiguration apiConfig) handlerReturnPostsFromFollowedFeeds(response http.ResponseWriter, request *http.Request, user database.User) {
	posts, err := apiConfiguration.database.GetPostsForUser(request.Context(), user.ID)

	if err != nil {
		errorJSON(response, 400, fmt.Sprintf("Não consegui retornar os seus posts: %v", err))
		return
	}

	JSONResponse(response, 200, posts)
}
