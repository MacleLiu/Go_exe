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

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	//确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	output := make(chan string)
	//并行遍历每一个文件树
	//fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		//go walkDir(root, &n, fileSizes)
		go func(root string) {
			defer n.Done()
			var w sync.WaitGroup
			fileSizes := make(chan int64)
			w.Add(1)
			go walkDir(root, &w, fileSizes)
			go func() {
				w.Wait()
				close(fileSizes)
			}()
			periodicOutput(root, output, fileSizes)
		}(root)
	}
	go func() {
		n.Wait()
		close(output)
	}()
	//定期输出结果
	//periodicOutput(fileSizes)
	for out := range output {
		fmt.Println(out)
	}
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema是一个用于限制目录并发数的计数信号量
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        //获取令牌
	defer func() { <-sema }() //释放令牌
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v", err)
		return nil
	}
	return entries
}

/* func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
} */

// 定期输出结果
func periodicOutput(root string, output chan string, fileSizes <-chan int64) {
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			//printDiskUsage(nfiles, nbytes)
			output <- fmt.Sprintf("%s\t%d files %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
		}
	}
	//printDiskUsage(nfiles, nbytes)
	output <- fmt.Sprintf("%s\t%d files %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}
