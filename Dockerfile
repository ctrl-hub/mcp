# Build stage
FROM golang:1.24-alpine AS builder

ARG VERSION=dev

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -extldflags '-static' -X main.version=${VERSION} \
    -a -installsuffix cgo \
    -o bin/ctrl-hub-mcp-server \
    cmd/server/main.go

# Final image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/bin/ctrl-hub-mcp-server /ctrl-hub-mcp-server

EXPOSE 8080

USER 65534:65534

ENTRYPOINT ["/ctrl-hub-mcp-server"]
