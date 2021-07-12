## GRPC使用

## 安装环境
- 安装protoc [链接](https://github.com/protocolbuffers/protobuf/releases)
  - 解压后将bin目录添加到$PATH中
- 安装protoc-gen-go
  - `go get -u github.com/golang/protobuf/protoc-gen-go`
  - 在$GOPATH外运行 会在$GOPATH/bin下生成二进制文件 同时需要将$GOPATH/bin添加到$PATH中
- `cd code-generate` 运行`bash gen-golang.sh` 生成`gym.pb.go`  
## 运行顺序
- 先运行server.go  再运行client.go
    
