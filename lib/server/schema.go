package server

import (
	"log"

	"github.com/graphql-go/graphql"
)

func Schema() graphql.Schema {
	schemaConfig := graphql.SchemaConfig{
		Query:    PdfQuery,
		Mutation: RenderMutation}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("[Server]", "Failed to create new schema, error: %v", err)
	}

	return schema
}
