FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o ./bin/main ./cmd/main.go

EXPOSE 3000

CMD ["/app/bin/main"]