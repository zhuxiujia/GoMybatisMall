#!/usr/bin/env bash
docker run -i --rm --name build_core --log-opt max-size=10m -v /var/lib/docker/volumes/jenkins_home/_data/workspace/build_core/:/go/src/github.com/zhuxiujia/GoMybatisMall/ -w /go/src/github.com/zhuxiujia/GoMybatisMall/core/main  golang:1.12.5 go build CoreService.go
