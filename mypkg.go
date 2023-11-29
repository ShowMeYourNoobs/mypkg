package main

import (
	"fmt"
	"mypkg/inner"
)

func main() {
	fmt.Println("mypkg")
	inner.InnerFunc()
}
