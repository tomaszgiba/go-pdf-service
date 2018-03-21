package server

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/tomaszgiba/gopdfservice/lib"
)

var pdfType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Pdf",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.Int,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var pageChannel = make(chan *lib.Page)

var pdfQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"get": &graphql.Field{
				Type:        pdfType,
				Description: "Retrive PDF by token",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					token, isOK := params.Args["token"].(string)
					var pdf (lib.Pdf)

					if isOK {
						fmt.Println(lib.PdfList[token])
						pdf = lib.PdfList[token]
					}

					return pdf, nil
				},
			},
		},
	})

var renderMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"render": &graphql.Field{
				Type:        pdfType,
				Description: "Render PDF",
				Args: graphql.FieldConfigArgument{
					"url": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					url, isOK := params.Args["url"].(string)
					var page (lib.Page)
					var pdf (lib.Pdf)

					if isOK {
						pdf.InitToken()
						page = lib.Page{URL: url, Body: nil}
						pdf.Page = &page
						lib.PdfList[pdf.Token] = pdf
						go page.DownloadBody(pageChannel)
					}

					return pdf, nil
				},
			},
		},
	})

func Schema() graphql.Schema {
	schemaConfig := graphql.SchemaConfig{
		Query:    pdfQuery,
		Mutation: renderMutation}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}
