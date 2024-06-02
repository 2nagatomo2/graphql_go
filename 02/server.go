package main

import (
	"log"
	"net/http"
	"os"

	"github.com/2nagatomo2/graphql_go/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// "/"をブラウザで開くとplaygroundが開く
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// "/query"にGraphQLリクエストを送ると結果が返ってくる
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
