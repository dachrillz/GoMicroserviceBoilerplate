#!/bin/bash

export CGO_ENABLED=0
cd accountservice;go get;go build -o accountservice-linux-amd64;echo built `pwd`;cd ..

docker build -t dachrillz/accountservice accountservice/
