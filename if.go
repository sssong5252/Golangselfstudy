package main

import "fmt"

func main() {
	var i int
	fmt.Print("num = ")
	fmt.Scanln(&i)
	fmt.Println("num is", i)

	if i == 4 {
		fmt.Print("the num is ", i)
	} else if i > 10{
		fmt.Println("too big")
	} else {
		fmt.Println("incorrect")
	}
}
