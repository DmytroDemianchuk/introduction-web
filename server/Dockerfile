FROM golang:1.21.1 AS builder

RUN go version

COPY . /github.com/Easy-Job-Developer/catalog_plus/
WORKDIR /github.com/Easy-Job-Developer/catalog_plus/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/Easy-Job-Developer/catalog_plus/.bin/app .

CMD ["./app"]