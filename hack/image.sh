#!/bin/bash

set -e

VERSION=$1
VERSION=${VERSION:-latest}
REGISTRY="registry.cn-hangzhou.aliyuncs.com"
NAMESPACE="huweihuang"
APP="zeus"
fullname=${REGISTRY}/${NAMESPACE}/${APP}:${VERSION}

function buildimage() {
    fullname=$1
    echo "docker build -t ${fullname} -f ./build/Dockerfile ."
    docker build -t "${fullname}" -f ./build/Dockerfile .
    echo "Building docker image ${fullname} succeeded."

    echo "docker push ${fullname}"
    docker push ${fullname}
    echo "Pushing docker image ${fullname} succeeded."
}

buildimage ${fullname}
