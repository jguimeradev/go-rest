FROM golang:1.26-alpine

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o ./bin/go-rest ./cmd

EXPOSE 6969

CMD ["bin/go-rest"]



