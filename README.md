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
go get github.com/EpochCloud/ConfCenter
注意这里报下面信息
package ConfCenter/initialization: unrecognized import path "ConfCenter/initialization" (import path does not begin with hostname)
package ConfCenter/router: unrecognized import path "ConfCenter/router" (import path does not begin with hostname)
这是github路径问题，不用管，直接进行如下
go get github.com/jmoiron/sqlx
go get github.com/BurntSushi/toml
mv ./src/github.com/EpochCloud/ConfCenter ./src/
mv ConfCenter ./src/ConfCenter/

添加日志路径和mysql地址
vim /src/ConfCenter/config/config.toml
根据要求填写

win环境
cd src/ConfCenter/
ConfCenter.exe -f ./config/config.toml
Linux/ MAC环境
cd src/ConfCenter/
./ConfCenter -f ./config/config.toml
```

## 数据库表结构

```
数据库的名字以及ip、端口、名字、密码都在toml配置文件中，如果使用本开源软件，请注意修改
表结构详情：在sql包中，如果想要测试，那么直接让sql文件导入navicat即可
```



