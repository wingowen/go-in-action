package main

import (
	"log"
	"os"

	_ "go-in-action/chapter2/sample/matchers"
	"go-in-action/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// 执行指定关键词的搜索
	search.Run("中国")
}
