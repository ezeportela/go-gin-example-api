FROM golang:1.16.4-alpine3.13

WORKDIR $GOPATH/src/github.com/ezeportela/meli-challenge

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["meli-challenge"]