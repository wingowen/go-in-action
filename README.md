# 环境配置

Windows 系统环境变量

- GOROOT：Go 安装路径（默认 C:\Go，安装程序自动设置）
- GOPATH：工作区路径（建议自定义，如 D:\go_workspace）
- PATH：添加 %GOROOT%\bin 和 %GOPATH%\bin

配置代理

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go env | findstr "GOPROXY"
```