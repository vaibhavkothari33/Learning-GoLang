// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Println("Taking input from user")

// 	reader:= bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter the number")
// 	input,_:=reader.ReadString('\n')

// 	fmt.Println("The multiplication of 4 with",input )

// 	multilpy,err:=strconv.ParseFloat(strings.TrimSpace(input),64)

// 	if err !=nil {
// 		fmt.Println(err)
// 	} else{
// 		fmt.Println("the multiplication is ",multilpy *4)
// 	}


// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Arrays in go lang")

// 	var names [5] string

// 	names[0] = "Vaibhav"
// 	names[1] = "Naman"
// 	names[3] = "Kothari"

// 	fmt.Println("Names list is",names)
// 	fmt.Println("the length of names is",len(names))

// }

package main

import "fmt"
func main() {
	var names = []string{"Vaibhav","Kothari","Naman","Jain"}

	fmt.Println("the names array is ",names)
	fmt.Printf("the names array %T ",names)

	names = append(names, "Madhu","Bapna")
	fmt.Println("")
	fmt.Println(" the new array is ",names)
}