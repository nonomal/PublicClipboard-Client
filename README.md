## PublicClipboard-Client

共享剪切板客户端程序 支持win、mac双系统

使用时需配置conf目录及conf.yml文件

### 相关仓库
[PublicClipboard-Server](https://github.com/sxz799/PublicClipboard-Server)

[PublicClipboard-Server-NoDB](https://github.com/sxz799/PublicClipboard-Server-NoDB)

[PublicClipboard-Web](https://github.com/sxz799/PublicClipboard-Web)

### 编译命令

#### mac/linux
`go build -ldflags="-s -w" -o Clipboard main.go`
#### windows
`go build -ldflags="-s -w" -o Clipboard.exe main.go`

