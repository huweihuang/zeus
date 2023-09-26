#!/bin/bash
set -x

version=$1

# 升级
kubectl set image deployment/zeus zeus=registry.cn-hangzhou.aliyuncs.com/huweihuang/zeus:${version} -n gin
# 查看滚动升级
kubectl rollout status deployment/zeus  -n gin
