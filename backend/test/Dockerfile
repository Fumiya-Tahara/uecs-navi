FROM golang:1.23-alpine3.19

WORKDIR /

RUN mkdir /backend

COPY ./backend ./backend

WORKDIR /backend

RUN apk update && apk add bash

RUN go mod download

RUN go install github.com/air-verse/air@latest

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

EXPOSE 8080

CMD [ "air", "-c", ".air.toml" ]
