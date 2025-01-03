FROM golang:alpine as builder

WORKDIR /go/src/github.com/cyber-xxm/gin-vue-admin
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"
# 设置时区
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata openntpd \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /go/src/github.com/cyber-xxm/gin-vue-admin

COPY --from=0 /go/src/github.com/cyber-xxm/gin-vue-admin/server ./
COPY --from=0 /go/src/github.com/cyber-xxm/gin-vue-admin/resource ./resource/
COPY --from=0 /go/src/github.com/cyber-xxm/gin-vue-admin/config.docker.yaml ./

# 挂载目录：如果使用了sqlite数据库，容器命令示例：docker run -d -v /宿主机路径/gva.db:/go/src/github.com/cyber-xxm/gin-vue-admin/gva.db -p 8888:8888 --name gva-server-v1 gva-server:1.0
# VOLUME ["/go/src/github.com/cyber-xxm/gin-vue-admin"]

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
