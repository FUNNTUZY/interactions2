PROTO_FILES=interactions.proto
PROTO_PATH=.

all: generate

generate:
	protoc \
	  --proto_path=$(PROTO_PATH) \
	  --proto_path=C:/Users/AlexK192/protovalidate/proto/protovalidate \
	  --proto_path=C:/Users/AlexK192/googleapis \
	  --go_out=$(PROTO_PATH) \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=$(PROTO_PATH) \
	  --go-grpc_opt=paths=source_relative \
	  --grpc-gateway_out=$(PROTO_PATH) \
	  --grpc-gateway_opt=paths=source_relative \
	  --validate_out="lang=go,paths=source_relative:$(PROTO_PATH)" \
	  $(PROTO_FILES)
