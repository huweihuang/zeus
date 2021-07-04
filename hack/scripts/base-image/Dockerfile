FROM alpine:latest
LABEL maintainer="Hu Weihuang"

# Set timezone
RUN apk update && apk add tzdata && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata
RUN apk upgrade && apk add --no-cache ca-certificates
RUN apk add --update tcpdump && rm -rf /var/cache/apk/*
ENV LANG='zh_CN.UTF-8'
