package main

import (
	"echo/config"
	"echo/db"
	"echo/server"
	"flag"
	"fmt"
	"os"
)

func main() {
	// 定义命令行参数 -e
	env := flag.String("e", "development", "environment")
	// 在使用编译后的文件时输出帮助信息: ./main -h
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	// 将命令行的参数解析为定义的标志
	flag.Parse()
	// 初始化环境配置
	config.Init(*env)
	// 初始化数据库
	db.Setup()
	// 启动服务
	server.Init()
}
