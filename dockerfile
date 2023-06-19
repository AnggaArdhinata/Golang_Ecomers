FROM golang:1.20.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

ENV DB_HOST=${DB_HOST} \
    DB_PORT=${DB_PORT} \
    DB_NAME=${DB_NAME} \
    DB_USER=${DB_USER} \
    DB_PASSWORD=${DB_PASS}

RUN go build -v -o /app/gocomers

EXPOSE 6625

ENTRYPOINT [ "/app/gocomers" ]
CMD [ "go run server.go" ]