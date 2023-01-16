# rpc配置proto文件生成rpc代码
# 进入rpc/pb目录
goctl rpc protoc user.proto --go_out=./ --go-grpc_out=./ --zrpc_out=../ -style goZero
