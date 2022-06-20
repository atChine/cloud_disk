# 这是一个云盘项目，使用Go-zero构建 实现简易的网盘系统

# 自动使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero


# 启动服务
go run core.go -f etc/core-api.yaml


可以将define.go 文件中的配置修改为自己的配置，然后启动服务