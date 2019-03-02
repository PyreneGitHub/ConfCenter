# README

## 自述

```go
ConfCenter是一个基于go语言开发微服务配置中心。
特点：
服务注册，服务发现
实现推拉两种功能。容错率低
轻量级
强大的并发能力
跨语言，跨平台
静态编译，轻松上云
所有语言都可以轻松使用
占用内存小
...
```



## 运行

```GO
cd $ GOPATH
 去获取github。com / EpochCloud / ConfCenter
cd src / github。COM / EpochCloud / ConfCenter
 去安装
cd $ GOPATH
mv ConfCenter $ GOPATH    //这里如果是windows是ConfCenter.exe
赢环境
ConfCenter。exe -f ./src/ConfCenter/config/config。toml
Linux的/ MAC环境
./ConfCenter -f ./src/ConfCenter/config/config。tomlcd GOPATH/ConfCenter
go run main.go -f ./config/config.toml
```

## 数据库表结构

```
数据库的名字以及ip、端口、名字、密码都在toml配置文件中，如果使用本开源软件，请注意修改
表结构详情：在sql包中
```



