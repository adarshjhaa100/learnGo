package main

import (
	"fmt"
	"math/rand"
	"rsc.io/quote" // import a module, go tidy will download this module if this is in a remote repo
)


func init() {
	rand.Seed(200)
}

func main() {
	fmt.Println("Hello, World")

	fmt.Println(quote.Opt())
}
