FROM golang:1.22.2

WORKDIR /app
COPY . .

RUN go mod download

COPY script/start-api.sh script/.

RUN chmod +x script/start-api.sh

CMD ["./script/start.sh"]