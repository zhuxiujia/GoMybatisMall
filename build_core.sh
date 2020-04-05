#!/usr/bin/env bash
docker run -i --rm --name build_core --log-opt max-size=10m -v /var/lib/docker/volumes/jenkins_home/_data/workspace/build_core/:/go/src/gitee.com/gascard/gas_server/ -w /go/src/gitee.com/gascard/gas_server/core/main  golang:1.12.5 go build CoreService.go
