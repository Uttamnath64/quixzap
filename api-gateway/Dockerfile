# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY services/auth ./services/auth
COPY app ./app
COPY scripts/start.sh ./scripts/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /quixzap-user ./cmd/server/main.go

# Production Stage
FROM alpine:3.18 AS production

WORKDIR /app

COPY --from=builder /quixzap-user /app/quixzap-user
COPY --from=builder /app/scripts/start.sh /app/scripts/start.sh

RUN chmod +x /app/quixzap-user /app/scripts/start.sh
RUN apk add --no-cache tzdata

EXPOSE 8082
CMD ["/app/scripts/start.sh"]
