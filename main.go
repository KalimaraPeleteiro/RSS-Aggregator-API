package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Carregando o Arquivo .env
	godotenv.Load()

	// Buscando a Porta que a aplicação irá executar.
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Não consegui acessar a porta em .env!")
	}

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

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Iniciando na Porta %v", port)
	err := server.ListenAndServe() // Esse comando irá parar e lidar com os requisitos até um erro ocorrer.

	if err != nil {
		log.Fatal(err)
	}
}
