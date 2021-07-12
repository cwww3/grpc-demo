#!/bin/bash

protoDir="../protos"
outDir="../languages/golang"
# -I： 指定import路径，可以指定多个-I参数，编译时按顺序查找，不指定默认当前目录
# -go_out：指定生成.pb.go文件的位置  --go_out=plugins=grpc: 生成grpc代码
protoc -I ${protoDir}/ ${protoDir}/*proto --go_out=plugins=grpc:${outDir}