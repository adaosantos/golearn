package main

import (
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"net/http"
	"golearn/modules"
	"golearn/helpers"
	"golearn/helpers/database"
	"fmt"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    modules.QueryType,
	Mutation: modules.MutationType,
})

func main() {
	port := helpers.GetENVorDefault("PORT", "9000")

	database.StartDB()
	defer database.CloseDB()

	// create a graphql-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	fmt.Print(http.ListenAndServe(":"+port, nil))
}