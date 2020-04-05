#!/usr/bin/env bash
docker stop consul
docker rm consul
docker run -d --net=host --log-opt max-size=10m --name=consul -p 8500:8500  -e CONSUL_BIND_INTERFACE=eth0  consul:1.4.0  agent  -dev  -client 0.0.0.0 -ui