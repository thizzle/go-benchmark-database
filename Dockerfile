FROM golang:1.4

RUN mkdir -p /go/src/github.com/thizzle/go-benchmark-database
WORKDIR /go/src/github.com/thizzle/go-benchmark-database

RUN go get bitbucket.org/liamstask/goose/cmd/goose github.com/jmoiron/sqlx

COPY . /go/src/github.com/thizzle/go-benchmark-database

CMD [ "go", "test", "-bench", ".", "-benchmem" ]
