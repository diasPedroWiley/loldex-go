FROM golang:1.19 as build-env

COPY go.mod go.sum /go/src/github.com/diasPedroWiley/loldex-go/
WORKDIR /go/src/github.com/diasPedroWiley/loldex-go
RUN go mod tidy
COPY . /go/src/github.com/diasPedroWiley/loldex-go

RUN go install github.com/githubnemo/CompileDaemon

EXPOSE 3003
ENTRYPOINT CGO_ENABLED=0 GOOS=linux CompileDaemon -build="go build -o /go/bin/loldex-go" -command="/go/bin/loldex-go"