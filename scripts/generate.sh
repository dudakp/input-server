#!/bin/bash

apppath=$(realpath ../../cmd/"$1")
protopath=$apppath/proto
protofile=$(realpath "$apppath"/proto/"$2")
grpcout="$apppath"/app/grpc

echo "generating grpc code to $grpcout from proto definition: $protofile"

protoc --go_out="$grpcout" --go_opt=paths=source_relative \
  --go-grpc_out="$grpcout" --go-grpc_opt=paths=source_relative \
  --proto_path="$protopath" "$protofile"
