FROM golang:latest

WORKDIR /var/www/go
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /prodtest_api

EXPOSE 8000

CMD ["/prodtest_api"]