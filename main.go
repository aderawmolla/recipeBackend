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

	r := routes.InitRoutes()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the API"))
	}).Methods("GET")
	log.Println("Server is running on port 2002...")
	log.Fatal(http.ListenAndServe(":2004", r))
}
