package main

import "fmt"

// create a type person struct
// attach method speak to type pointer to a person
// *person
// create a type human interface
// create func "saysomething"
// have it take in a human as a perameter
// have it call the speak method
// show the following in your code
// you CAN pass *person into saysomething
// you CANNOT pass person into saysomething

type person struct{
	first string
}

type human interface{
	speak()
}

func (p *person) speak(){
	fmt.Println("hello")
}

func saysomething(h human){
	h.speak()
}

func main(){

	p1:= person{
		first: "ali",
	}
    // does not work
	// saysomething(p1)
	saysomething(&p1)
	p1.speak()


}