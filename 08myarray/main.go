package main

import "fmt"

func main() {
	fmt.Println("The arrays in go lang")

	var fruitlist [4] string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Mango"
	// fruitlist[2] = ""
	fruitlist[3] = "Banana"

	fmt.Println("Fruit list is ", fruitlist) // Fruit list is  [Apple Mango (space) Banana]

	fmt.Println("The length of the fruitlist is ",len(fruitlist)) // 4 but i only have 3 items

	var vegList = [3] string{"Potato","Beans","Onion"}
	
	fmt.Println("Veg List is ",len(vegList))
	fmt.Println("Veg List is ",vegList)
	// fmt.Println("Veg List is ",vegList[5]) // out of bound
	fmt.Println("Veg List is ",vegList[2])
}