# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY scripts/start.sh scripts/.
COPY scripts/start-consumer.sh scripts/.
# RUN chmod +x scripts/start.sh scripts/start-consumer.sh
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /quick-connect-app ./cmd/server/main.go
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /quick-connect-consumer ./cmd/consumer/main.go

# Production stage (server)
FROM golang:1.23-alpine AS production
WORKDIR /app
COPY --from=builder /quick-connect-app /app/quick-connect-app
COPY --from=builder /app /app
COPY --from=builder /app/scripts/start.sh /app/scripts/start.sh
RUN chmod +x /app/quick-connect-app /app/scripts/start.sh
COPY internal/app/templates/ /app/internal/app/templates/
RUN apk add --no-cache tzdata
EXPOSE 8081
CMD ["/app/scripts/start.sh"]

# Consumer stage
# FROM golang:1.23-alpine AS consumer
# WORKDIR /app
# COPY --from=builder /quick-connect-consumer /app/quick-connect-consumer
# COPY --from=builder /app /app
# COPY --from=builder /app/scripts/start-consumer.sh /app/scripts/start-consumer.sh
# RUN chmod +x /app/quick-connect-consumer /app/scripts/start-consumer.sh
# RUN apk add --no-cache tzdata
# CMD ["/app/scripts/start-consumer.sh"]