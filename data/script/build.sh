#!/bin/bash

root_path=/data/golang/src/word-game-go
proto_path=$root_path/data/protobuf/proto
out_path=$root_path/data/rpc
platform=golang

if [ ! -d $out_path/$platform ]; then
    mkdir -p $out_path/$platform
fi

for file in $proto_path/*
do
    protoc --go_out=$out_path/$platform $file
done
