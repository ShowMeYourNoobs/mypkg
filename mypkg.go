package main

import (
	"fmt"
	"mypkg/inner"
)

func CallInner() {
	fmt.Println("Calling inner func")
	inner.InnerFunc()
}

func CallFunc() {
	fmt.Println("Main func")
}
