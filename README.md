# restful-api-demo

go语言的 业务纵向分割模板

我之前的项目是mvc项目, 内部耦合性较大
现在从业务出发, 将模块拆分为纵向, 一棍子捅到底, 一个业务就是一个项目
对之前的项目进行拆分

```
cmd:CLI
apps:业务

```


# 环境
1. protoc [去下载](https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.1)
2. protoc-gen-go [v1.31.0]
3. protoc-gen-go-grpc [1.3.0]

```cmd
# 安装指定版本的方式
go install k8s.io/klog@v1.0.0
```
