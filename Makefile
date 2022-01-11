# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=api-service
BINARY_UNIX=$(BINARY_NAME)_unix

#ejecuta la compilación con el modo verbose
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# ejecuta test
test:
	$(GOTEST) -v ./...

#limpia los binarios compilados
clean:
	$(GOCLEAN)
	rm -f $(BINARY_UNIX)
# ejecuta el main.go
run:
	$(GORUN) main.go

run-bin:
#$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v