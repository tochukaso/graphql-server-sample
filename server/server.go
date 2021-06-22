package main

import (
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
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

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ProductObservers: map[string]chan *model.Product{},
	}}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

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
