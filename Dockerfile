FROM golang:1.16-alpine as builder
WORKDIR /GoProject/src/github.com/yanhuan0802/uc
ENV GOPATH=/GoProject

# Copy the go source
COPY go.mod go.mod
COPY httpserver/ httpserver/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -o bin/httpserver ./httpserver/http_server.go

FROM alpine:3.12
WORKDIR /
LABEL maintainers="huan.yan"
ENV VERSION=1.0.0

COPY --from=builder /GoProject/src/github.com/yanhuan0802/uc/bin/httpserver .
RUN chmod +x /httpserver

EXPOSE 80
ENTRYPOINT ["/httpserver"]