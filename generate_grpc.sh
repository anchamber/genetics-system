#! /bin/bash
# protoc --go_out=. --go-grpc_out=. api/proto/ping.proto
protoc --go_out=. --go-grpc_out=. --proto_path=../genetics-api/proto system.proto
protoc --go_out=. --go-grpc_out=. --proto_path=../genetics-api/proto api.proto