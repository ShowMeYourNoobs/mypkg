package mypkg

import (
	"fmt"
	"mypkg/inner"
)

func MyPkg() {
	fmt.Println("mypkg")
	inner.InnerFunc()
}

