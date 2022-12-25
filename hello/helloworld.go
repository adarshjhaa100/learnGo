package main 

import (
		"fmt"
		strUtility "example/hello/morestring" // import package morestring inside of hello and give alias strUtility
		gogcmp "github.com/google/go-cmp/cmp" //downloading external package(always import package, not module)
		) 

func main() {
	fmt.Println("Hello World!!!")

	fmt.Println("Reverse of string: abcde", strUtility.ReverseString("abcde"))
	fmt.Println("Concatenate abc with defg", strUtility.ConcatString("abc", "defg"))
	fmt.Println("Comparison of two string: ", gogcmp.Diff("Hello", "World"))
}
