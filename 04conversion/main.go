// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {

// 	fmt.Println("Welcome to pizza app")
// 	fmt.Println("Please rate out pizza from 1 to 5")

// 	reader := bufio.NewReader(os.Stdin)

// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("Thanks for rating")
// 	fmt.Println("input", input)
// 	// fmt.Println("input", input+1) // string can not add 1

// 	// numrating:= input + 1
// 	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("added one ", numrating+1)
// 	}
// }
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
	fmt.Println("enter the number between 1 and 5")

	reader:=bufio.NewReader(os.Stdin)

	input,_:=reader.ReadString('\n')
	fmt.Println("the number is ",input)

	// addition,err:=strconv.ParseFloat(strings.TrimSpace(input),64)

	addition,err:=strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("the addition is ",addition +5)
	}

}