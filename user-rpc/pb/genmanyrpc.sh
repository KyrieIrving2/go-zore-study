#！bin/sh

protoc -I ./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. order.proto