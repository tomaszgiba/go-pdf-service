package server

import (
	"github.com/graphql-go/graphql"
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
					var page (Page)
					var pdf (Pdf)

					if isOK {
						pdf.Init()

						go func() {
							// get this out of here
							page = Page{URL: url, Body: nil, FilePath: TempFilePath(pdf.Token)}
							pdf.Page = &page
							PdfList[pdf.Token] = pdf
							DownloadPageBody(&pdf)
							SavePageToFile(&pdf)
							RenderAndSavePdf(&pdf)
							go UploadPdfToS3(&pdf)
							pdf.Finalize()
						}()

					}

					return pdf, nil
				},
			},
		},
	})
