FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM debian:bookworm-slim

RUN useradd -m appuser

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY keys /app/keys
RUN chown -R appuser:appuser /app

USER appuser
EXPOSE 8080

ENTRYPOINT ["./main"]
