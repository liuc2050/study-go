package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: error %v\n", err)
				continue
			}

			if dup, _ := countLines(f, counts); dup == true {
				fmt.Printf("file name:%s\n", f.Name())
			}

			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) (bool, error) {
	var ret bool
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++

		if counts[input.Text()] > 1 {
			ret = true
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dup2: error %v\n", err)
		return ret, err
	}

	return ret, nil
}
