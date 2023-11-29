package main

import (
	"fmt"
	"mypkg/inner"
)

var shouldPrint2 bool = false

func main() {
	fmt.Println("mypkg")
	inner.InnerFunc()

	shouldPrint1 := false
	if shouldPrint1 {
		fmt.Println("1")
	}

	if shouldPrint2 {
		fmt.Println("2")
		fmt.Println("hola")
	}
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")

	fmt.Println("start loop")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
