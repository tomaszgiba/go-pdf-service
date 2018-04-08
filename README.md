# go-pdf-service

run: `go run main.go`

**Goal**: to have PDF microservice, that is easily deployable to AWS, fast and easy to set up.


### to send PDF to render:

```
curl -XPOST http://localhost:8080/graphql/pdf/ \
	-H 'Content-Type: application/graphql' \
	-d 'mutation { render(url: "https://golang.org/pkg/time/"){url, token, state} }'
```

or: 

```
curl -g 'http://localhost:8080/graphql/pdf/?query=mutation+_{render(url:"https://golang.org/pkg/time/"){token}}'
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