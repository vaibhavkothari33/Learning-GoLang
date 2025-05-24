package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	greeting:="hello world"
	fmt.Println("Enter your name")
	reader:=bufio.NewReader(os.Stdin)

	input,err:=reader.ReadString('\n')
	fmt.Println("welcom",input)
	fmt.Println("error",err)
	fmt.Println(greeting)
}