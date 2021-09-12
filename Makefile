.PHONY: pb gen

all: pb gen build

pb:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/storagepb/storage.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/trackerpb/tracker.proto

gen:
	go build -o bin/speedfs-gen tools/gen/*.go
	bin/speedfs-gen -input proto/storage/storage.go -rpc 1
	bin/speedfs-gen -input proto/tracker/tracker.go -rpc 1
	bin/speedfs-gen -input proto/cmd.go
	bin/speedfs-gen -input proto/header.go

build:
	go build github.com/speedfs/speedfs/proto
	go build github.com/speedfs/speedfs/proto/storage
	go build github.com/speedfs/speedfs/proto/tracker