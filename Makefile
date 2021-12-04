.PHONY: pb rpcgen

RPCGEN=bin/speedfs-rpcgen

all: pb rpcgen build test

pb:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/storagepb/storage.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/trackerpb/tracker.proto

rpcgen:
	go build -o ${RPCGEN} tools/rpcgen/*.go
	${RPCGEN} -input proto/storage/storage.go -rpc 1
	${RPCGEN} -input proto/tracker/tracker.go -rpc 1
	${RPCGEN} -input rpc/ping.go
	${RPCGEN} -input rpc/quit.go
	${RPCGEN} -input rpc/reply.go
	${RPCGEN} -input rpc/header.go

build: pkg cmd

pkg:
	go build github.com/speedfs/speedfs/internal/storage
	go build github.com/speedfs/speedfs/proto
	go build github.com/speedfs/speedfs/proto/storage
	go build github.com/speedfs/speedfs/proto/tracker
	go build github.com/speedfs/speedfs/rpc

cmd: cli

cli: client
	go build -o bin/speedfs-cli github.com/speedfs/speedfs/cmd/cli

client:
	go build github.com/speedfs/speedfs/client

test:
	go test github.com/speedfs/speedfs/internal/storage

clean:
	rm -vf bin/*