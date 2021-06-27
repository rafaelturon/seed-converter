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

### Certificates
```shell
openssl genrsa -des3 -out rootCA.key 2048 # Create Root signing Key
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1825 -out rootCA.pem #Generate self-signed Root certificate

openssl genrsa -out client.key 2048 # Create a Key certificate for the Server
openssl req -new -key client.key -out client.csr # Create a signing CSR
openssl x509 -req -in client.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out client.crt -days 825 -sha256 # Generate a certificate for the Server
```