FROM golang:1.18.3

WORKDIR app

COPY cmd/client ./cmd/client
COPY internal/general ./internal/general
COPY internal/client ./internal/client
COPY pkg ./pkg
COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go build -o /word-of-wisdom-client ./cmd/client

CMD ["/word-of-wisdom-client"]