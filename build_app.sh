#!/usr/bin/env bash
# build app
docker run -i --rm --name build_app --log-opt max-size=10m -v /var/lib/docker/volumes/jenkins_home/_data/workspace/build_app/:/go/src/gitee.com/gascard/gas_server/ -w /go/src/gitee.com/gascard/gas_server/app/main  golang:1.12.5 go build App.go
