package main

import (
	"fmt"
	"os"
)

// 打印每个参数的索引和值，每个一行
func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, " ", arg)
	}

}
