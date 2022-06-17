# 自动使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero


# 启动服务
go run core.go -f etc/core-api.yaml
