package main

import (
	_ "./matchers"
	"./search"
	"log"
	"os"
)

func init() {
	// 将入职输出到标准输入
	log.SetOutput(os.Stdout)
}

func main() {
	// 使用特定的项做搜索
	search.Run("president")
}

