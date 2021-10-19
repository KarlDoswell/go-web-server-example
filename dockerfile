FROM golang:1.16-alpine

WORKDIR /myApp

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /example-web-server

EXPOSE 8080

CMD [ "/example-web-server" ]
