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

## time control

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println("Formated Time")
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) 

	createdDate:=time.Date(2020,time.January,19,23,23,0,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	// date time week
}
```