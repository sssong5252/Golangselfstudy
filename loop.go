package main

import "fmt"

func main() {
	var i int
	fmt.Scanln(&i)
	var j string
	var h int
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("i am 3")
		} else {
			fmt.Println("hello")
		}
		fmt.Scanln(&j)
		if j == "0" {
			fmt.Println("the end")
		}
		fmt.Println("hello")
	}
	fmt.Println("while start")
	for {
		fmt.Print("insert the key")
		fmt.Scanln(&h)
		if h == 0 {
			break
		} else {
			fmt.Println("hello")
		}
	}
}
