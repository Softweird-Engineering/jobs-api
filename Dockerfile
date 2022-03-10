FROM golang:buster
ARG GOPROXY=https://goproxy.cn,direct

COPY ./src /app
WORKDIR /app
RUN go mod download
RUN go build -o main

ENTRYPOINT [ "./main" ]
