APP_NAME = api-go

build:
	docker build -t $(APP_NAME):dev .

run: build
	docker run -it -p 80:80 --rm --name $(APP_NAME) $(APP_NAME):dev

shell:
	docker exec -it $(APP_NAME) /bin/sh

lint:
	go fmt ./...

static:
	go vet ./...

test: lint static
	go test ./...
