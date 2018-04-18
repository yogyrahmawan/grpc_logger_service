PROJECT_DIR="${GOPATH}/src/github.com/yogyrahmawan/logger_service"

proto:
	protoc -I src/proto/ -I "${GOPATH}/src/" --go_out=plugins=grpc:"${GOPATH}/src" src/proto/logger.proto src/proto/logger_response.proto src/proto/logger_responses.proto	
	protoc -I src/proto/ -I "${GOPATH}/src/" -I "${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" --go_out=plugins=grpc:"${GOPATH}/src" src/proto/logger_service.proto

clean_proto:
	rm -Rf src/api/*.pb.go
	rm -Rf src/domain/*.pb.go
