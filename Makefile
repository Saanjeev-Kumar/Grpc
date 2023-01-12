create:
	protoc --proto_path=hellopb hellopb/*.proto --go_out=./
	protoc --proto_path=hellopb hellopb/*.proto --go-grpc_out=./

clean:
	rm gen/proto/*.go