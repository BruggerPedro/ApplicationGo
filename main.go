package main

import (
	"ApplicationGo/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// @title           API de gerenciamento de estoque
// @version         1.0
// @description     Uma API que oferece o gerenciamento de estoque de pecas e materiais.
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	fmt.Println("Conexão feita!")
	routes.LoadRoutes()
	http.ListenAndServe(":"+port, nil)
}
