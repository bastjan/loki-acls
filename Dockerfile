FROM golang:1

WORKDIR /app

COPY . .

RUN go build

CMD ["/app/loki-acls"]
