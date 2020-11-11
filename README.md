# grpc-gateway-example

## Run server by following command

The gRPC server listens on `tcp port 50050`, and Http reverse proxy server listens on `tcp port 8080`.

```
go run main.go server.go struct.go
```

## Test functionality

```
curl -X POST localhost:8080/employee/create -d '{"Name": "William", "Age": 24}'
```
