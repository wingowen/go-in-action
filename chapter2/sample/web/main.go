package main

import (
	"log"
	api "go-in-action/chapter2/sample/web/api"
)

// main 是API服务器的入口点
func main() {
	log.Println("启动API服务器...")
	api.StartServer()
}