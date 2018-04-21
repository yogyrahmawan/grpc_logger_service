FROM yogyrahmawan/docker-go-debian

RUN apt-get -y install git curl && \ 
	mkdir -p /root/pathgo/src/github.com/yogyrahmawan/grpc_logger_service && \
	mkdir -p /root/pathgo/src/github.com/yogyrahmawan/cmd && \
	mkdir -p /root/pathgo/src/github.com/yogyrahmawan/src && \
        mkdir -p /root/cert

RUN mkdir -p $GOPATH/bin && \
	mkdir -p $GOPATH/src && \
	mkdir -p $GOPATH/pkg

WORKDIR /root/pathgo/src/github.com/yogyrahmawan/
COPY . grpc_logger_service/ 
COPY ./cmd/* cmd/
COPY ./src/* src/
COPY ./cert/* /root/cert/

WORKDIR /root/pathgo/src/github.com/yogyrahmawan/grpc_logger_service/
SHELL ["/bin/bash", "-lc"]
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only
RUN go build
RUN go install
