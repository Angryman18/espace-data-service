package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Angryman18/espace-data-service/client"
	Env "github.com/Angryman18/espace-data-service/env"
	"github.com/Angryman18/espace-data-service/graph"
)

const defaultPort = "8080"

func main() {
	env := Env.GetEnv()
	var port string = env.Port
	if env.Port == "" {
		port = defaultPort
	}

	database, mongoClient := client.ConnectMongo(env.MongoURI)

	defer func() {
		mongoClient.Disconnect(context.TODO())
	}()

	newDb := graph.NewDB(database)
	newClient := client.NewClient(newDb.Database)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Client: newClient}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
