package main

import (
	"fmt"
	"mypkg/inner"
)

func main() {
	fmt.Println("mypkg")
	inner.InnerFunc()

	shouldPrint1 := false
	if shouldPrint1 {
		fmt.Println("1")
	}

	fmt.Println("2")
	fmt.Println("3")

	fmt.Println("start loop")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
