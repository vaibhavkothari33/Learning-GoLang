package main

import "fmt"

func main() {

	fmt.Println("hello 1")
	defer fmt.Println("World 1")
	defer fmt.Println("World 2")
	defer fmt.Println("World 3")
	defer fmt.Println("World 4")
	fmt.Println("hello 2")
	mydefer() 
	// defer alway execute in the last 
	// (LIFO)
	// Last in first out
	// ello1
	// hello2
	// Worl4
	// Worl3
	// Worl2
	// Worl1

	
}

func mydefer(){
	for i :=0; i<5; i++{
		defer fmt.Println("hello from function ",i)
	}
}