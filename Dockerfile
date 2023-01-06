FROM golang:1.19-alpine

WORKDIR /app/api
ENV GOBIN=/usr/local/go/bin

RUN apk update && apk add --no-cache \
    curl bash wget

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -mod=mod -o /github.com/diasPedroWiley/loldex-go

EXPOSE 3003

# CMD [ "/github.com/diasPedroWiley/loldex-go" ]
# CMD CompileDaemon -command="./github.com/diasPedroWiley/loldex-go"
CMD sh entrypoint.sh