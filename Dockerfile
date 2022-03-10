FROM golang:buster

COPY ./src /app
WORKDIR /app
RUN go mod download
RUN go build -o main

ENTRYPOINT [ "./main" ]
