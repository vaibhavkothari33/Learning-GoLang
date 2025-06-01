package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
func main() {
	fmt.Println("Files in golang")

	content:= "this needs to go in a file"

	file,err:=os.Create("./mygofile.txt")

	if err !=nil {
		fmt.Println("Error")
		panic(err)
	}

	length,err := io.WriteString(file,content)

	// if err !=nil{
	// 	fmt.Println("error")
	// 	panic(err)
	// }
	checkNilError(err)
	fmt.Println("length is ",length)
	defer file.Close()
	readFile("./mygofile.txt")
	
}

func readFile(filename string){
	databyte,err:= ioutil.ReadFile(filename)

	checkNilError(err)
	fmt.Println("text data inside the file is \n",string(databyte))


}

func checkNilError(err error){
	if err !=nil{
		panic(err)
	}
}