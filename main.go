package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	database *database.Queries
}

func main() {

	// Carregando o Arquivo .env
	godotenv.Load()

	// Buscando Variáveis de Ambiente
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Não consegui acessar a porta em .env!")
	}

	databaseURL := os.Getenv("DB_URL")

	if databaseURL == "" {
		log.Fatal("Não consegui achei o banco de dados em .env!")
	}

	// Biblioteca nativa de SQL do Go
	connection, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Não consegui conectar ao banco de dados.")
	}

	database := database.New(connection)
	apiConfiguration := apiConfig{
		database: database,
	}

	go startScraping(database, 10, time.Minute)

	// Criando o Servidor
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter() // Paginação
	v1Router.Get("/", handlerServerReadiness)
	v1Router.Get("/error", handlerDefaultError)
	v1Router.Post("/users/create", apiConfiguration.handlerCreateUser)
	v1Router.Get("/users/getUser", apiConfiguration.middlewareAuth(apiConfiguration.handlerGetUserByAPIKey))
	v1Router.Post("/feeds/add", apiConfiguration.middlewareAuth(apiConfiguration.handlerCreateFeed))
	v1Router.Get("/feeds/all", apiConfiguration.handlerGetAllFeeds)
	v1Router.Post("/users/follow", apiConfiguration.middlewareAuth(apiConfiguration.handlerFollowNewFeed))
	v1Router.Get("/users/my_feeds", apiConfiguration.middlewareAuth(apiConfiguration.handlerGetAllUserFollowingFeeds))
	v1Router.Delete("/users/unfollow/{feed_id}", apiConfiguration.middlewareAuth(apiConfiguration.handlerUnfollowFeed))
	v1Router.Get("/users/my_feeds/posts", apiConfiguration.middlewareAuth(apiConfiguration.handlerReturnPostsFromFollowedFeeds))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Iniciando na Porta %v", port)
	err = server.ListenAndServe() // Esse comando irá parar e lidar com os requisitos até um erro ocorrer.

	if err != nil {
		log.Fatal(err)
	}
}
