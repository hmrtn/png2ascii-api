FROM golang:1.16.2

WORKDIR /app

COPY ./app .

RUN go mod download

RUN go build -o /server

EXPOSE 8080
