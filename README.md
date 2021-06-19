# Seed Converter


## Testing
### Run Tests
```shell
go test ./...
```
### Cover Tests
```shell
go test ./... -cover
```
### Cover Tests Output
```shell
go test ./... -coverprofile=coverage.out
```
### Display HTML Output
```shell
go tool cover -html=coverage.out
```

## Documentation
### Godoc
```shell
godoc -http=:6060
```
### Protoc
```shell
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
protoc --go_out=. --go-grpc_out=. entity.proto
```