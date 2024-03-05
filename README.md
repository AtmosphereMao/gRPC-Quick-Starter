# gRPC Quick Starter
## Prerequisites
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Usage
### Proto
 - src/proto (gRPC Example by https://github.com/grpc/grpc-go)
 - src/service (Custom Proto)
```shell
$ protoc -I . --go_out=./src/service --go_opt=module="grpc-server/src/service" --go-grpc_out=./src/service --go-grpc_opt=module="grpc-server/src/service" ./src/service/service.proto
```

### Run
#### gRPC Example
```shell
go run .\src\greeter_server\main.go
go run .\src\greeter_client\main.go
```
#### gRPC Hello World
```shell
go run .\src\job_server\server.go
go run .\src\job_client\client.go
```
#### gRPC Stream
```shell
go run .\src\stream_server\server.go
go run .\src\stream_client\cilent.go
```
#### Auth (Middleware authentication)
```shell
go run .\src\auth_server\server.go
go run .\src\auth_client\client.go
```

### Reference

[grpc.io](https://grpc.io/docs/languages/go/quickstart/)

[gRPC 详细入门](https://cloud.tencent.com/developer/article/2266206)