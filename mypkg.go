package main

import (
	"fmt"

	"github.com/ShowMeYourNoobs/mypkg/inner"
)

func CallInner() {
	fmt.Println("Calling inner func v1.0.0")
	inner.InnerFunc()
}

func CallFunc() {
	fmt.Println("Main func  v1.0.0")
}
