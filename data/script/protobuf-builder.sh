#!/bin/bash

root_path=/data/golang/src/word-game-go
proto_path=$root_path/data/protobuf/proto
out_path=$root_path/data/rpc
platform=golang

## 判断目录是否存在
if [ ! -d $out_path"/"$platform ];then
    mkdir -p $out_path"/"$platform
else
    for file in `ls $out_path"/"$platform`; do
        rm -rf $out_path"/"$platform"/"$file
    done
fi

## 解析指定目录下的所有proto文件
for file in `ls $proto_path`; do
    if [ "${file##*.}" = "proto" ]; then
        if [ ! -d $out_path"/"$platform"/"${file%%.*} ]; then
            mkdir $out_path"/"$platform"/"${file%%.*}
        fi
        protoc --proto_path=$proto_path --go_out=$out_path"/"$platform"/"${file%%.*} $proto_path"/"$file
    fi
done
