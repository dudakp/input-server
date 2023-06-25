#!/bin/bash

app_name=$1
app_path=$(realpath ../../cmd/"$app_name")

commons_path=$(realpath ../../internal/proto)

protofile_path=$(realpath "$app_path"/proto)
protofile=$protofile_path/$2

grpc_out=$app_path/app/infrastructure

echo "starting generating grpc code for app: $app_name"

echo "generating common grpc code from proto definition: $commons_path"
protoc --go_out="$commons_path" \
  --go_opt=paths=source_relative \
  --proto_path="$commons_path" \
  --proto_path="$protofile_path" \
  "$commons_path"/*.proto
echo "done generating common code"

echo "generating grpc code to $grpc_out from proto definition: $protofile"

protoc --go_out="$grpc_out" \
  --go_opt=paths=source_relative \
  --proto_path="$commons_path" \
  --proto_path="$protofile_path" \
  --go-grpc_out="$grpc_out" \
  --go-grpc_opt=paths=source_relative \
  "$protofile"

echo "done generating grpc code"
