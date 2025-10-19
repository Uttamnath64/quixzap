# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY services/widget ./services/widget
COPY app ./app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /quixzap-widget ./cmd/server/main.go

# Production Stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /quixzap-widget /app/quixzap-widget

RUN apk add --no-cache tzdata

CMD ["/app/quixzap-widget"]
