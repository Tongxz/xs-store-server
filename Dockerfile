FROM golang:alpine as builder

WORKDIR /go/src/github.com/Tongxz/xs-store-server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o xs-server .

FROM alpine:latest

LABEL MAINTAINER="bluestonea@foxmail.com"

WORKDIR /go/src/github.com/Tongxz/xs-admin-vue/server

COPY --from=0 /go/src/github.com/Tongxz/xs-store-server/xs-server ./
COPY --from=0 /go/src/github.com/Tongxz/xs-store-server/resource ./resource/
COPY --from=0 /go/src/github.com/Tongxz/xs-store-server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./xs-server -c config.docker.yaml
