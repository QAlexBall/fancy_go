package main

import (
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
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