package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/inadati/gqlkit"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("GRAPHQL_SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// router.Use(responseHeader.Middleware())
	// router.Use(auth.Middleware())

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(gqlkit.NewExecutableSchema(gqlkit.Config{Resolvers: &gqlkit.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Printf("graphql endpoint is http://localhost:%s/query", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
