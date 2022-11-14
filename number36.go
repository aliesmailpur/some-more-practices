package main

import (
	"fmt"
	
)

// write a progran that
// launches 10 goroutines
// each goroutine add 10 numbers to a channel
// pull the numbers off the channel and print them

func main(){
	c:= make(chan int)


    for j:=0;j<10;j++{
		go func(){
			for i:=0;i<10;i++{
				c<-i
			}
			
		}()

	}

	for k:=0;k<100;k++{
		fmt.Println(<-c)
	}

	fmt.Println("about to exit")
	
}