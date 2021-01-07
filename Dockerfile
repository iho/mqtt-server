FROM golang:1.14

WORKDIR /go/src/app
RUN go get -v github.com/rubenv/sql-migrate/...

COPY . .
COPY docker-conf.json conf.json

RUN go build cmd/server/main.go
RUN sql-migrate up
CMD ["./main"]