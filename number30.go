package main

import ( 
	"fmt"
	"runtime"
	"sync"
)

// Using goroutines, create an incrementer program
// have a variable to hold the incrementer value
// lunch a bunch of goroutine
// each goroutine should
// read incrementer value 
// store it in a new variable
// yield the processor with gosched()
// increment the new variable
// with the value in the variable back ti the incrementer variable

func main(){

	fmt.Println(runtime.NumCPU())

	var wg sync.WaitGroup

	
    

	counter := 0

	const gs = 50

	wg.Add(gs)


	for i:=0;i<gs;i++{
		
		go func(){
			v:= counter
			
			v++
			
			counter = v
			wg.Done()

		}()

	

	}
	wg.Wait()
	fmt.Println(counter)



}
