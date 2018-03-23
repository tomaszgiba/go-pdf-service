package main

import (
	"fmt"
	"net/http"

	gqlhandler "github.com/graphql-go/handler"
	s "github.com/tomaszgiba/gopdfservice/lib/server"
)

func main() {
	fmt.Println("[Server]", "Starting")
	renderSchema := s.Schema()
	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &renderSchema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql/pdf/", handler)

	fmt.Println("[Server]", "Started")

	// and serve!
	http.ListenAndServe(":8080", nil)
}
