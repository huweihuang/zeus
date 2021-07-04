#!/bin/bash

set -x
set -e

REGISTRY="registry.cn-hangzhou.aliyuncs.com"
NAMESPACE="huweihuang"
SERVICE="alpine"
VERSION="base"

docker build --network=host -t ${REGISTRY}/${NAMESPACE}/${SERVICE}:${VERSION} ./

docker push ${REGISTRY}/${NAMESPACE}/${SERVICE}:${VERSION}
