package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cnt := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		cnt[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Println("input error!")
		return
	}

	fmt.Println("Now print your dumplicate lines:")

	for line, n := range cnt {
		if n > 1 {
			fmt.Printf("%d %s\n", n, line)
		}
	}

	fmt.Println("Test map iterator:")
	for line, n := range cnt {
		if n > 1 {
			fmt.Printf("%d %s\n", n, line)
		}
	}
}
