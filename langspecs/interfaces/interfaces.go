package interfaces

import (
	"fmt"
)

// interface is defined as type that has set of method signatures
// for a variable to be assigned to an interface var,
// it must implement all the methods defined in the interface

type Abser interface {
	Abs() float64
}

type Vertex struct {
	x, y float64
}

type SampleName struct {
	s string
}

type Empty interface {}

// this method implements interface, but we don't explicitly need to 
// specify that explicitly (like we do in java using implements keyword)
func (v *Vertex) Abs() float64{
	return (v.x*v.x + v.y*v.y)
}

func (s * SampleName) Abs() float64{
	return 1.1
}


func InterfaceTest() {
	var intrf Abser
	v := Vertex {2,3}

	intrf = &v // we should have atleast those methods implemented by 
				// the type which are declared by the interface

	fmt.Println(intrf.Abs()) //13
	
	// Under the hood, and interface value(here intrf) can be thought of
	// as a tuple that holds a value and its concrete type (value, type)
	// Here it is: (&{2 3}, *interfaces.Vertex)
	
	fmt.Printf("(%v, %T, %#v)\n", intrf, intrf, intrf)


	// an interface value can even point to null, 
	// the method call here will be valid although with a nil receiver arg(not nullpointer exception)
	var s *SampleName
	intrf = s	
	fmt.Println(intrf.Abs())
	fmt.Printf("(%v, %T, %#v)\n", intrf, intrf, intrf)
	// fmt.Printf("\n(%#T)", intrf)


	// Empty Interface (can hold value of any type)
	fmt.Println("\nEmpty interface can switch to values of any type")
	
	var emp Empty
	emp = 1
	fmt.Println(emp)
	emp = 2.2
	fmt.Println(emp)
	emp = Vertex {1,2}
	fmt.Println(emp)


	// Type assertion: check if interface  holds a specific type
	fmt.Println("Type Assertions")
	emp = 1
	fmt.Println(emp.(int))
	emp = 2.2
	fmt.Println(emp.(float64))
	emp = Vertex {1,2}
	// fmt.Println(emp.(float64)) // here emp is a vertex, so we'll have a panic (panic: built in go functionality that stops the normal flow, defer works though)

	val1, ok1 := emp.(float64)
	val2, ok2 := emp.(Vertex)

	// to test for tyoe
	fmt.Println("Testing types")
	fmt.Println(val1, ok1) // on a wrong type assertion
	fmt.Println(val2, ok2) // on a correct type assertion


	str := ""
	for i := 0 ; i<10 ; i++{
		str+="dws"
	}
	fmt.Println(str)

	// stringer : defined in fmt, useful where we have to return string (the string may contain any types in the end will return string)
	// Its generally implemented using Sprintf
	fmt.Printf("\n%v\n",fmt.Sprintf("%v %v", 3, 3.4))


}
