package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	fmt.Println("Another style:")

	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}
