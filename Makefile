PROJECT_DIR="${GOPATH}/src/github.com/yogyrahmawan/logger_service"

proto:	
	protoc -I src/proto/ -I "${GOPATH}/src/" -I "${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" --go_out=plugins=grpc:"${GOPATH}/src/github.com/yogyrahmawan/logger_service/src/pb" src/proto/*

proto_gateway:
	protoc -I src/proto/ -I "${GOPATH}/src/" -I "${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" --grpc-gateway_out=logtostderr=true:"${GOPATH}/src/github.com/yogyrahmawan/logger_service/src/pb" src/proto/logger_service.proto

clean_proto:
	rm -Rf src/pb/*.pb.go

gen_proto_moc:
	mockgen -destination=src/mockspb/mocks_proto.go -package=mockspb github.com/yogyrahmawan/logger_service/src/pb LoggerServiceClient

generate_cert:
	rm -Rf cert/*
	openssl genrsa -out cert/server.key 2048
	openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
	openssl req -new -sha256 -key cert/server.key -out cert/server.csr
	openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650

run_docker: 
	docker-compose up 

test: 
	go test ./... 

build: 
	go build 

.PHONY proto proto_gateway clean_proto gen_proto_moc generate_cert run_docker test build
