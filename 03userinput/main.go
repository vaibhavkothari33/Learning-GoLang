// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// // comma ok syntax

// func main() {
// 	welcome := "welcome to go"
// 	fmt.Println(welcome)
// 	reader:= bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter rating for the pizza")
// 	//comma okay syntax

// 	// error ok 
// 	// input,_:=reader.ReadString('\n')
// 	input,err:=reader.ReadString('\n')
// 	fmt.Println("thanks for typing", input)
// 	fmt.Printf("thanks for typing %T", input)
// 	fmt.Println("")
// 	fmt.Println("thanks for typing ", err)
// }


package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	welcome:="welcome to go "
	fmt.Println(welcome)
	fmt.Println("Enter your name")
	reader:=bufio.NewReader(os.Stdin)

	input,err:=reader.ReadString('\n')
	fmt.Println("the name of the user is: ",input)
	fmt.Printf("the Type of the name is: %T",input)
	fmt.Println("")
	fmt.Println("the name of the user is: ",err)
}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("taking input from the user")

// 	reader:=bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter you name")

// 	input,err:=reader.ReadString('\n')

// 	fmt.Println("welcome" ,input)
// 	fmt.Println("the error is",err)
// }