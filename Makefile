gen:
	protoc --proto_path=pkg/proto --go_out=./pkg/pb --go-grpc_out=./pkg/pb pkg/proto/*.proto

run-server:
	go run cmd/main.go

.PHONY:
	gen run-server