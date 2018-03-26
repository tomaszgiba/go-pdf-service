# go-pdf-service

run: `go run main.go`


### to send PDF to render:

```
curl -XPOST http://localhost:8080/graphql/pdf/ \
	-H 'Content-Type: application/graphql' \
	-d 'mutation { render(url: "https://golang.org/pkg/time/"){url, token, state} }'
```
and you will a token

### to get one PDF:

```
curl -XPOST http://localhost:8080/graphql/pdf/ \
		-H 'Content-Type: application/graphql' \
		-d 'query { get(token: "HRNqIxkhBTKL"){url, token, state} }'
```

### to get all PDFs:

```
curl -XPOST http://localhost:8080/graphql/pdf/ \
		-H 'Content-Type: application/graphql' \
		-d 'query { all{url, token, state} }'
```