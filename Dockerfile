FROM golang:1.14

RUN mkdir /go/src/hangul-api
WORKDIR /go/src/hangul-api
COPY go.mod .
COPY go.sum .
COPY cmd ./cmd

RUN go build -o /bin/hangul-api ./cmd/server

ENTRYPOINT ["/bin/hangul-api"]
