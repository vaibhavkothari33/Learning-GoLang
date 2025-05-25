package main

import "fmt"

func main() {
	fmt.Println("Pointers in go lang")

	// var ptr* int
	// fmt.Println("The value of pointer is",ptr)
	
	myNumber:= 23

	var ptr = &myNumber
	fmt.Println("the value of ptr is ",ptr) // )xc0000008798
	fmt.Println("the actual value of ptr is ",*ptr) // 23
	*ptr = *ptr *2
	fmt.Println("the new value is",*ptr) // 46
}