protoc -I proto/ --go_out=src/ proto/simple.proto
protoc -I=proto/ --go_out=src/ proto/simple.proto