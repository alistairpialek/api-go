APP_NAME = api-go
GIT_COMMIT = $(shell git rev-parse HEAD)

build:
	docker build --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(APP_NAME):dev .

run: build
	docker run -it -p 80:80 --rm --name $(APP_NAME) $(APP_NAME):dev

shell:
	docker exec -it $(APP_NAME) /bin/sh

lint:
	go fmt ./...
	golangci-lint run ./...

static:
	go vet ./...

test: lint static
	go test -v ./...

install:
	brew install golangci-lint
