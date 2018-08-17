package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress meassages")
var sema = make(chan struct{}, 1000) //limit concurrency

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//Traverse the file tree.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	//Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
	var end bool
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				end = true
				break
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
		if end {
			break
		}
	}
	printDiskUsage(nfiles, nbytes)
}

//dfs
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        //acquire token
	defer func() { <-sema }() //release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}

	//remove symbol links
	end := 0
	for i, entry := range entries {
		if entry.Size() > 1e9 {
			fmt.Printf("name:%s/%s size:%d\n", dir, entry.Name(), entry.Size())
		}
		if entry.Mode()&os.ModeSymlink != 0 {
			//skip
		} else if dir == "/proc" && entry.Name() == "kcore" {
			//not real file, skip
		} else {
			if i != end {
				entries[end] = entry
			}
			end++
		}
	}
	return entries[:end]
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
