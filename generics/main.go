package main


import (
	"fmt"
	types "learning/generics/generictype"
)


func main() {
	fmt.Printf("Hello Generics!!!\n")

	lst := types.List[string]{}	
	fmt.Printf("\nInitial List: %#v\n", lst)


	a := lst.Insert("Node1", -1)
	a = lst.Insert("Node2", -1)
	a = lst.Insert("Node3", -1)
	a = lst.Insert("Node1.1", 1)

	if(a == 0) {
		lst.Traverse()
	}
}