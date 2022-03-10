FROM golang:buster

COPY ./src/go.mod /app/
COPY ./src/go.sum /app/
WORKDIR /app

RUN go mod download

COPY ./src /app
RUN go build -o main

ENTRYPOINT [ "./main" ]
