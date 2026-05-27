daemon:
	go run cmd/ardaemon/main.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/autorun.proto

sqlc:
	sqlc generate