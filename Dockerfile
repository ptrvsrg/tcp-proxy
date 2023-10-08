FROM golang:1.18

WORKDIR /tcp-proxy

COPY go.mod go.sum ./
COPY cmd ./cmd
COPY internal ./internal

EXPOSE $PROXY_PORT

RUN go mod download
ENTRYPOINT go run ./cmd/tcp-proxy -port $PROXY_PORT