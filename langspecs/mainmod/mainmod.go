package main

// import modules
import (
	"fmt"
	"math/cmplx"
	// "learn/controlstatements"
	"learn/additionalDS" 
)

// functions: takes more than one args and can return more than one args
func divide(p int, q int) (int, int) {
	return p / q, p % q
}

// can omit the type from all but last, others will
func divide1(p, q int) (int, int) {
	return p / q, p % q
}

// function: named return value, return statement returns empty
// dont use this shit
func namedReturn(sum int) (x, y int) {
	x = sum + 1
	y = sum - 1

	return
}

// declare variables
func declareVars() {
	// declaring variables,
	// as function params, variables takes on type of righmost defined
	var isPresent, isNew, isHull bool

	// intialize browsers
	var a, b int = 2, 3

	// short declarations
	call, draw, full := 1, "hello", true

	fmt.Println(isPresent, isNew, isHull)
	fmt.Println(a, b)
	fmt.Println(call, draw, full)

	

	// Constants, cannot be declared using a:= syntax
	const I int = 3824723
	const PI float32 = 3.1416

	// untyped const takes value according to its context
	const (
		EUREKA = 1 << 100
		REUKA  = EUREKA >> 99
	)

	fmt.Println(float64(EUREKA))
	fmt.Println(REUKA)

}

// basic types
func basicTypes() {
	// bool
	// string (No char present, can convert int to char)
	// signed: int, int8/16/32/64, unsigned: uint8/16/32/64, rune; same as int32
	// byte: same as int8
	// float: float32/64
	// complex: complex64/128
	// intptr: intptr, uintptr

	// clubbing mutiples variables inside blocks "()"
	var (
		StringVar  string     = "HELLO_STRING"
		MaxInt     uint64     = 1<<64 - 1
		ComplexVar complex128 = cmplx.Sqrt(-2 + 3i)
	)

	fmt.Println(StringVar, MaxInt, ComplexVar)

	// Printf: analogous to c printf
	// use %T for type
	fmt.Printf("Type: %T for %v\n", ComplexVar, ComplexVar)

	// values declared without their initial values are assigned their default 0s
	// numeric: 0, bool: false, string: ""

}

// Explicit conversion: there is no explicit conversion in go
func explicitConversion() {
	a, b, c := 23, 43.2, uint(2)

	fmt.Println(float64(a))
	fmt.Println(int64(b))
	fmt.Println(float64(c))
}


// we can pass more than one value to function using ...
func moreThanOne(listvar ...any) {
	
	// any: stands for interface {}
	fmt.Println("List of any:", listvar) //prints a list of items
}



func main() {

	fmt.Println("Hello!")

	// arithmetic: same as in c (int/int = int, int/float any combination gives float)

	// fmt.Println(4.0/3)

	// fmt.Println(divide(78, 23))

	// fmt.Println(namedReturn(50))

	// declareVars()

	// basicTypes()

	// i, j := "helo", 45

	// fmt.Printf("\n%T : %v, %T : %v\n", i,i,j,j)

	// explicitConversion()

	// ----- control statements -----

	// controlstatements.RunForLoop(10)

	// controlstatements.ForWhile(10)

	// controlstatements.RunForever()

	// controlstatements.IfStatementTest()

	// controlstatements.Sqrt(3)

	// controlstatements.SwitchStatement()

	// controlstatements.DeferTest()

	// ----- control statements ----

	// ----- other data structures: pointers,structs,maps,slices

	// ----- other data structures: pointers,structs,maps,slices
	// additionalDS.TestPkg()
	// additionalDS.PointerType()
	// additionalDS.StructTest()
	// additionalDS.ArrayTypes()
	// additionalDS.Slices()
	// additionalDS.RangeTest()
	// moreThanOne(2,3,4,"dsd",4.5)
	// additionalDS.MapTest()
	// additionalDS.WordCountMap()
	// additionalDS.PassFunction()
	additionalDS.ClosureTest()

}
