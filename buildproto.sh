export PATH="$PATH:$(go env GOPATH)/bin"
protoc --proto_path=proto/ --go_out=plugins=grpc:pb --go_opt=paths=source_relative proto/*.proto
