package main

import "fmt"

func add(){
	fmt.Printf("This is an addition function\n");
	a:= 64;
	b:= 36;
	c:= a+b;
	fmt.Printf("Value after addition is %d \n",c);
}
func main(){
	fmt.Printf("Hello Mahesh welcome to Go-lang !! \n");
	add();
}
