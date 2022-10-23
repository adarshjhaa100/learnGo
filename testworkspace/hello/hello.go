package main // the main package define as main, other package with package names

import (
	"fmt"
	"wrktest/calculate"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Addition of 2 numbers: ", calculate.Add2Num(100,200) )
}