# Build the Go binary
FROM golang:1.24.2 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/web/server

RUN go build -o /goapi

# Final stage
FROM debian:bookworm-slim

COPY --from=builder /goapi /goapi

EXPOSE 3001

CMD ["/goapi"]
