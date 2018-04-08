package server

import (
	"github.com/graphql-go/graphql"
	"github.com/tomaszgiba/go-pdf-service/lib/model"
)

var RenderMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"render": &graphql.Field{
				Type:        PdfType,
				Description: "Render PDF",
				Args: graphql.FieldConfigArgument{
					"url": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"expires_in": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var isOK = true
					url, isOK := params.Args["url"].(string)
					expiresIn, isOK := params.Args["expires_in"].(int)
					var page (model.Page)
					var pdf (model.Pdf)

					if isOK {
						page.Init(url, nil)
						pdf.Init(&page, expiresIn)
						go func() {
							pdf.DownloadPageBody()
							pdf.SavePageToFile()
							pdf.RenderAndSavePdf()
							go pdf.UploadPdfToS3() // WARN: first delegate upload, then set expire
							pdf.Finalize()
						}()
					}

					return pdf, nil
				},
			},
		},
	})
