package main

import "fmt"

func main() {
	var i int
	fmt.Scanln(&i)

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("i am 3")
		}
		fmt.Println("hello")
	}
}
