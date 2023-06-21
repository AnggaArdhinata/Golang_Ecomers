FROM golang:1.20.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -v -o /app/gocomers

EXPOSE 6625

ENTRYPOINT [ "/app/gocomers" ]
CMD [ "go run server.go" ]