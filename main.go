package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang/start/go/config"
	"golang/start/go/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	graphqlURL := os.Getenv("HASURA_GRAPHQL_URL")
	fmt.Println(graphqlURL)
	services.InitGraphQLClient(graphqlURL)

	r := routes.InitRoutes() // Initialize routes

	log.Println("Server is running on port 2003...")
	log.Fatal(http.ListenAndServe(":2003", r))
}
