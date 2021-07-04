#!/bin/bash
set -x

version=$1

# 升级
kubectl set image deployment/gin-api-frame gin-api-frame=registry.cn-hangzhou.aliyuncs.com/huweihuang/gin-api-frame:${version} -n gin
# 查看滚动升级
kubectl rollout status deployment/gin-api-frame  -n gin
