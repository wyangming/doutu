package main

import (
	_ "doutu/module"
	"fmt"
	//"runtime"
	"sync"
)

func main() {
	//cpu密集型项目时充分利用cpu性能
	//runtime.GOMAXPROCS(runtime.NumCPU())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("a")
	wg.Wait()
}
