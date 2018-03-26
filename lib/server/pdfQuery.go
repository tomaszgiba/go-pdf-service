package server

import (
	"github.com/graphql-go/graphql"
)

var PdfQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"get": &graphql.Field{
				Type:        PdfType,
				Description: "Retrive PDF by token",
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					token, isOK := params.Args["token"].(string)
					var pdf (Pdf)

					if isOK {
						pdf = PdfList[token]
					}

					return pdf, nil
				},
			},
			"all": &graphql.Field{
				Type:        graphql.NewList(PdfType),
				Description: "List of all PDFs",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return PdfList, nil
				},
			},
		},
	})
