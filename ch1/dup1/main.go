package main

import (
	"bufio"
	"fmt"
	"os"
)

// 查找重复行 map bufio
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	// 遍历 map 键line 值n
	for line, n := range counts {
		// 书上是 > 1
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
