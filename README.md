# chan

> chan （蝉）是一个Web服务器，用于提供Web服务、反向代理

## 功能模块

+ [x] 静态资源服务
+ [x] 配置模板封装
    + [x] 基础配置
    + [x] SSL配置
    + [x] 支持多域名的TLS服务
+ [x] 反向代理
+ [x] 负载均衡
    + [x] 轮询
    + [x] 加权轮询
+ [ ] 命令行
    + [x] version 获取版本
    + [x] run 运行
    + [ ] start 启动
    + [ ] stop 停止
    + [ ] restart 重启
    + [ ] conf 获取配置文件路径
+ [ ] 服务运行

## 快速上手

+ 运行
```shell
# 不指定配置文件，默认在80端口启动一个服务
chan run
# 使用指定的配置文件路径启动Web服务
chan run -t 配置文件路径
chan run -t ./conf/chan.json
```