### 静态文件转发代理工具

#### 1. 安装
```shell
go install github.com/kingparks/goproxy@latest
```

#### 2. 使用 
##### -p 指定监听端口 -s 指定静态文件目录
```shell
goproxy -p :8081 -s ~/domains/public_html
```