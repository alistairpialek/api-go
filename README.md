# api-go

Sample Golang REST API using [Gorilla Mux](https://github.com/gorilla/mux) with unit tests.

## Requirements

* Docker
* Golang 1.19 for running tests

## Running

The API docker container can be run as a foreground process by running:

```
make run
```

## Endpoints

### /health

Provides a way to check if the application is healthy.

| Name        | Method       | Params | Example Response          |
| ----------- | ------------ | -------| ------------------------- |
| /health     | GET          | None   | 200 HTTP status code only. |

Example cURL:

```
curl -v http://0.0.0.0/health
```

### /metadata

Provides information about the running application.

| Name        | Method       | Params | Example Response          |
| ----------- | ------------ | -------| ------------------------- |
| /metadata   | GET          | None   | `{"git_commit":"abc123","invocations":1}` |

Example cURL:

```
curl -v http://0.0.0.0/metadata
```

### /calculate

Takes a request body with CPU and memory metrics for services and provides suggestions on
request and limits settings.

| Name        | Method       | Example Body | Example Response          |
| ----------- | ------------ | -------------| ------------------------- |
| /calculate  | POST         | `[{"app": "app","time": "20227-21T06:35:00.000Z","cpu usage (mcores)": 1.00,"memory usage (MiB)": 1.00}]` | `{"services":[{"name":"app","cpu":{"request":1,"limit":2},"memory":{"request":1,"limit":2}}]}` |

Example cURL:

```
curl -X POST -H "Content-Type: application/json" -d @metrics-input.json "http://0.0.0.0/calculate"
```

## Testing

Tests (including formatting and linting) can be run with:

```
make install (optional: to install golangci-lint)
make test
```
