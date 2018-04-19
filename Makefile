PROJECT_DIR="${GOPATH}/src/github.com/yogyrahmawan/logger_service"

proto:	
	protoc -I src/proto/ -I "${GOPATH}/src/" -I "${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" --go_out=plugins=grpc:"${GOPATH}/src/github.com/yogyrahmawan/logger_service/src/pb" src/proto/*

clean_proto:
	rm -Rf src/pb/*.pb.go
