FROM golang:buster

COPY ./src/go.mod /app/
COPY ./src/go.sum /app/
WORKDIR /app

RUN go mod download

COPY ./src /app
RUN if [ "$GIN_MODE" != "test" ] ; then go build -o main; fi

CMD [ "./main" ]


