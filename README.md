# go-xserver
Go 服务器框架（go-x.v2）

现在看 go-x ，制作时没有从使用者角度来做。因此重新写个。

## 目标

go-xserver 致力于实现 1 个高可用、高易用的 Golang 服务器框架

go-xserver 以插件的方式，来丰富框架内容

## 编译

- 编译环境需要[翻墙设置](doc/编译-翻墙设置.md)

- 编译执行以下语句即可：

  ```shell
  ./make.sh
  ```

- Windows 10 下开发，请参考[在Win10中Linux环境搭建](doc/编译-在Win10中Linux环境搭建.md)

## 已完成功能

- [主体框架](doc/规范-代码框架.md)
- [配置模块](doc/规范-配置文件.md)
- [服务发现](doc/框架层功能-服务发现.md)

## 将要实现的功能

- 框架层功能
    - 服务器组内互联
    - 服务器组内通信


- 逻辑层功能
    - 登陆服务
    - 网关服务
    - 大厅服务
    - 匹配服务
    - 房间服务
    - 管理服务
