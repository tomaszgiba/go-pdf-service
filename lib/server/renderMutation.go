package server

import (
	"github.com/graphql-go/graphql"
	"github.com/tomaszgiba/gopdfservice/lib"
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
					var page (lib.Page)
					var pdf (lib.Pdf)

					if isOK {
						pdf.InitToken()

						go func() {
							// get this out of here

							page = lib.Page{URL: url, Body: nil, FilePath: lib.TempFilePath(pdf.Token)}
							pdf.Page = &page
							lib.PdfList[pdf.Token] = pdf
							lib.DownloadPageBody(&pdf)
							lib.SavePageToFile(&pdf)
							lib.RenderAndSavePdf(&pdf)
							lib.UploadPdfToS3(&pdf)
							pdf.SignalReady()
						}()

					}

					return pdf, nil
				},
			},
		},
	})
