package main

import (
	"fmt"
	"os"
)

// 打印os.Args[0]，即被执行命令本身的名字
func main() {
	var name = os.Args[0]

	fmt.Println(name)
}
