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

# Memory allocation and deallocation happes automatically in go lang
```go
new ()

allocate memory but no init 
you will get a memory address

zeroes storage
```

```go
make()
Allocate memory and init 
you will get a memory address
non zeroes storage
```

garbage collection happed automatically
out of scope or nil

## Pointers in Go Lang

passing the memory address not the variable
actual value reference
& ===> reference

```go
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
```