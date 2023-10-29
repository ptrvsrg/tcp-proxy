FROM golang:1.18-alpine as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./cmd ./cmd
COPY ./internal ./internal
RUN go build -o /tcp-proxy ./cmd/tcp-proxy/main.go

FROM alpine:3
COPY --from=builder tcp-proxy /bin/tcp-proxy
EXPOSE ${PROXY_PORT}
ENTRYPOINT /bin/tcp-proxy -port ${PROXY_PORT}