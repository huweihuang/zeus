FROM registry.cn-hangzhou.aliyuncs.com/huweihuang/alpine:base
LABEL maintainer="Hu Weihuang"

COPY ./hack/conf/conf.toml /conf.toml
COPY ./_output/bin/gin-api-frame /gin-api-frame

CMD ["/gin-api-frame"]
