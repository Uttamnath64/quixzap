FROM golang:1.23.5

WORKDIR /app
COPY . .

RUN go mod download

COPY scripts/start.sh scripts/.

RUN chmod +x scripts/start.sh

CMD ["./scripts/start.sh"]