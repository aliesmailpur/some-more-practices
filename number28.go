package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// in addiction to the main goroutine, launch two additional goroutines
// each addtional goroutine should print something out
// use waitgroups

func main(){
	fmt.Println(runtime.NumGoroutine())
    wg.Add(2)
	go ali()

	go foo()

	bar()

	wg.Wait()

}

func ali(){
	for c:=0;c<10;c++{
		fmt.Println("ali" , c)
	}
	wg.Done()
}

func foo(){
	for i:=0;i<5;i++{
		fmt.Println("foo" ,i)
	}
	wg.Done()
}

func bar(){
	for v:=0;v<5;v++{
		fmt.Println("bar : " ,v)
	}
}