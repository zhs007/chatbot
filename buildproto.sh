export PATH="$PATH:$(go env GOPATH)/bin"
protoc --proto_path=protos/ --go_out=plugins=grpc:chatbotpb --go_opt=paths=source_relative protos/*.proto
protoc --proto_path=basicchatbotprotos/ --go_out=plugins=grpc:basicchatbotpb --go_opt=paths=source_relative basicchatbotprotos/*.proto
