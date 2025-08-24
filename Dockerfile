FROM golang:1.23.5

WORKDIR /app
COPY . .

RUN go mod download

COPY start-api.sh .

RUN chmod +x start-api.sh

CMD ["./start.sh"]