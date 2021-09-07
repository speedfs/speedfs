.PHONY: pb gen

pb:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/storagepb/storage.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/trackerpb/tracker.proto

gen:
	go build -o bin/speedfs-gen tools/gen/*.go
	bin/speedfs-gen -input proto/storage/storage.go -output proto/storage/storage.proto.go
	bin/speedfs-gen -input proto/tracker/tracker.go -output proto/tracker/tracker.proto.go