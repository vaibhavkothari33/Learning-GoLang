## how to take the input
```go
package main

import "fmt"

func main(){
    fmt.Println("Enter your name")
    reader:=bufio.NewReader(os.Stdin)

    input,_:=reader.ReadString('\n')
    fmt.Println("your name is ",input)

}
```

## how to do the conversion
```go
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
```