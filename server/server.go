package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"github.com/tochukaso/graphql-server/db"
	"github.com/tochukaso/graphql-server/env"
	"github.com/tochukaso/graphql-server/graph"
	"github.com/tochukaso/graphql-server/graph/generated"
	"github.com/tochukaso/graphql-server/graph/model"
)

const defaultPort = "8083"

func main() {
	createSchemas()

	port := env.GetEnv().Port
	if port == "" {
		port = defaultPort
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	router := http.NewServeMux()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, graph.LoaderMiddleware(router)))
}

func createSchemas() {
	db := db.GetDB()
	db.AutoMigrate(
		&model.Product{},
		&model.Sku{},
	)
}
