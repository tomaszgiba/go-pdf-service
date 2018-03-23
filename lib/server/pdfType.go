package server

import "github.com/graphql-go/graphql"

var PdfType = graphql.NewObject(
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
