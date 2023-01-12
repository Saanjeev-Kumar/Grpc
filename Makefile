create:
	protoc --proto_path=proto/user --go_out=proto/pb --go_opt=paths=source_relative \
        --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative \
         --grpc-gateway_out=proto/pb --grpc-gateway_opt=paths=source_relative \
        proto/user/user.proto


