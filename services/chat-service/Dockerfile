# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY services/chat ./services/chat
COPY app ./app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /quixzap-chat ./cmd/server/main.go

# Production Stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /quixzap-chat /app/quixzap-chat

RUN apk add --no-cache tzdata

CMD ["/app/quixzap-chat"]
