package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		// 文件打开错误跳过
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3 %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%s\t%d", line, n)
		}
	}

}
