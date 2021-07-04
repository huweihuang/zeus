#!/bin/bash

set -x

option=$1
name=$2

url="127.0.0.1:8080"

case ${option} in
"create")
    curl -X POST -H 'Content-Type: application/json' -d@${option}.json $url/api/v1/instance |jq
    ;;
"delete")
    curl -X DELETE -H 'Content-Type: application/json' -d@${option}.json $url/api/v1/instance |jq
    ;;
"get")
    curl -X GET $url/api/v1/instance?name=${name} |jq
    ;;
*)
    echo "option not found"
    ;;
esac
