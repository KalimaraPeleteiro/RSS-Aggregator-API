package main

import (
	"fmt"
	"log"
	"os"

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

	fmt.Println("Iniciando na Porta", port)
}
