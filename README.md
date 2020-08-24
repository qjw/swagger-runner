## 准备
+ 安装依赖 `go mod tidy`
+ 打包静态资源 `go generate`
+ 编译 `go build .`
+ 静态编译 `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./swagger-runner .`

## 运行

``` bash
$ ./swagger-runner -h
Usage of ./main:
  -port uint
    	bind port. (default 12345)
  -spec string
    	spec file path. (default "spec.json")
```
