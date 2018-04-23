#!/bin/sh

echo "Waiting for mongodb connection"

while ! nc -z mongodb 27017; do 
	sleep 0.1 
done 

echo "Mongodb started"

cd /root/go/bin/ && ./grpc_logger_service --config /root/pathgo/src/github.com/yogyrahmawan/grpc_logger_service/config.toml --env test

