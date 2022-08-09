FROM golang:1.19.0-bullseye

ARG LISTEN_PORT=80
ENV LISTEN_PORT=${LISTEN_PORT}

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change.
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app

# Required by github.com/gorilla/mux (request router)
EXPOSE ${LISTEN_PORT}

CMD ["app"]
