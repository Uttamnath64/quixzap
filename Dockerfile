FROM golang:1.22.2

WORKDIR /app
COPY . .

RUN go mod download

COPY start-api.sh .

RUN chmod +x start-api.sh

CMD ["./start.sh"]