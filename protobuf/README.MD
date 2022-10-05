
# Protocol Buffers简介
> 现在的网络应用都是前后端分离的，数据传输方式有：json和xml两种格式，其中json更多一些。现在又多了一种数据传输方式，就是google开发的Protocol Buffers。在分布式应用或者微服务架构中，各个服务之间通常使用json或者xml结构数据进行通信，通常情况下，是没什么问题的，但是在高性能和大数据通信的系统当中，如果有办法可以压缩数据量，提高传输效率，显然会给用户带来更快更流畅的体验。

### 也就是现在的数据传输方式有三种：

* JSON
* xml
* Protocol Buffers
---

## 下载安装Protocol Buffers编译器

### 下载地址
```
https://github.com/protocolbuffers/protobuf
```

## vscode扩展
```
vscode-proto3
```

## 创建第一个protobuf文件

```
syntax = "proto3";  // 必须的，表示使用proto3

option go_package="./;hello"; // 表示编译go文件存放地址和包名

package hello;

message Person{
  string name = 1;
  int32 age = 2;
  string email = 3;
  enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 4;
}
```

### 安装go包
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### 编译命令
```
protoc --go_out=./go ./proto/*
```


### 报错：'protoc-gen-go' 不是内部或外部命令，也不是可运行的程序或批处理文件。
* 查看 GOPATH 的 bin 是否有相关文件
* 将相关文件复制到 GOROOT 的 bin 目录下 