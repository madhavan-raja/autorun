run:
	go run .

proto:
	protoc --go_out=. --go-grpc_out=. proto/autorun.proto