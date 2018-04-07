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
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					url, isOK := params.Args["url"].(string)
					var page (model.Page)
					var pdf (model.Pdf)

					if isOK {
						pdf.Init()

						go func() {
							// get this out of here
							page = model.Page{URL: url, Body: nil, FilePath: model.TempFilePath(pdf.Token)}
							pdf.Page = &page
							model.PdfList[pdf.Token] = pdf
							model.DownloadPageBody(&pdf)
							model.SavePageToFile(&pdf)
							model.RenderAndSavePdf(&pdf)
							go model.UploadPdfToS3(&pdf)
							pdf.Finalize()
						}()

					}

					return pdf, nil
				},
			},
		},
	})
