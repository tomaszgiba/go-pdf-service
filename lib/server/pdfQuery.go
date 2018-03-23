package server

import (
	"github.com/graphql-go/graphql"
	"github.com/tomaszgiba/gopdfservice/lib"
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
					var pdf (lib.Pdf)

					if isOK {
						pdf = lib.PdfList[token]
					}

					return pdf, nil
				},
			},
		},
	})
