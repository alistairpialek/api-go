FROM golang:1.19.0-bullseye as builder

ARG GIT_COMMIT

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change.
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v \
  -ldflags="-X 'github.com/alistairpialek/api-go/v1/utils.GitCommit=${GIT_COMMIT}'" \
  -o /usr/local/bin/app

FROM golang:1.19.0-bullseye as final

ARG LISTEN_PORT=80
ENV LISTEN_PORT=${LISTEN_PORT}

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

# Required by github.com/gorilla/mux (request router)
EXPOSE ${LISTEN_PORT}

CMD ["app"]
