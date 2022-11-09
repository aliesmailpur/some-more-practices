package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// use package ATOMIC to prevent the race condition



func main(){

	fmt.Println(runtime.NumCPU())

	var wg sync.WaitGroup

	

	
    

	var counter int64

	const gs = 50

	wg.Add(gs)


	for i:=0;i<gs;i++{
		
		go func(){

			atomic.AddInt64(&counter, 1)
			fmt.Println( atomic.LoadInt64(&counter))	
					
			
			wg.Done()

		}()

	

	}
	wg.Wait()
	fmt.Println(counter)



}
