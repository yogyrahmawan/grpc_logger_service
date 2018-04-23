# grpc_logger_service
Mini centralized log based on grpc 

### Clone 
```
git clone https://github.com/yogyrahmawan/grpc_logger_service.git 
```

### Run using docker 
```
docker-compose up 
```

### Run locally 
```
go build 
export jwt_token="token"
./grpc_logger_service --config path/to/config --env test/development/production 
```

### Run test 
```
go test ./...
```

### Configuration 
configuration is located on file config.toml 

### Short explanation 
1. database mongodb 
2. protocol : grpc and http 1 for get logging 
3. Using jwt for auth 
4. see make file about how to generate protocol buffer, mock, etc
etc
