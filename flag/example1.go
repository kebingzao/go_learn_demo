package main

import (
	"fmt"
	"flag"
)


func main() {
	// 定义一个命令行参数 port，并设置默认值为 9088
	port := flag.String("port", "9088", "service port")
	// 解析命令行参数
	flag.Parse()
	// 最后输出
	fmt.Println("port is ", *port)
}