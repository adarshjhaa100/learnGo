package controlstatements

import (
	"fmt"
	"runtime"
)


func RunForLoop(n int) {
	sum:= 0

	// no parenthesis surrounding the control statements
	// the init and post statements are optional
	for i:=0; i<=n; i++ {
		sum+=i	
	}
	
	fmt.Println("Sum:", sum)
}

func ForWhile(n int) {
	sum:= 0
	// we can drop semicolons here
	for n>=0 {
		sum += n
		n--
	}

	fmt.Println("Sum using for-while:", sum)
}

// runs forever, nothing much to explain
func RunForever() {
	
	for { 
		fmt.Println("going on forever") 
		
	}
}

func IfStatementTest() {
	// if statement doesnt require parenthesis
	x := 10 // stands for var x: int = 10

	if x%5 == 0 {
		fmt.Println("x is a multiple of 10")
	}

	// short statements (like for, we can use short statements like the init). 
	// The scope of vars decalred within the short statement is local to the for
	// ** also, else should start just after if scope braces ends(} )
	if v := 31; v%6 == 0 {
		fmt.Printf("%v is a multiple of 6.\n", v);
	} else {
		fmt.Printf("%v is a not multiple of 6.\n", v);
	}
	// can't use v here, out of scope

}

func Sqrt(x float64) {
	z := 1.0
	z = float64(1)

	fmt.Println("Z at each stage")

	// Newton's method: z = z - (z*z - x)/2*x 

	for i := 0; i < 10; i++ {
		z = z - (z*z - x)/(10*z)
		fmt.Println(z)
	} 
}

func SwitchStatement() {
	fmt.Println("actual OS: ", runtime.GOOS)

	// short statements can be used in switch
	// break is automatically provided in go.
	// evaluation order: top to bottom

	switch os := runtime.GOOS; os {
		case "darwin": 
			fmt.Println("OS X")
		case "linux":
			fmt.Println("linux")
		default:
			fmt.Printf("%s \n", os)
	}

	// switch without no condition (can be used to write long if else statements)
}

func DeferTest() {
	// defer will evaluate the parameters but will defer the 
	// statement until alll surrounding statements are evaluated
	// multiple defers: here, the first defer will have the defer on 
	// line wid Statement 1 will wait for the statement following it to execute (even if its other defer) 
	// defer statement stack up (like an actual stack really, allowing other statement to execute and only defers are pushed to stack)
	// Application: we can defer statements which close resources just after opening them (files, connections etc.). 
	// This will ensure that they will run at the end.
	// Even if the surrounding function exits, defer will run, its pushed on a separate call stack of itself

	// defer becomes useful for cleanup and stuff 
	fmt.Println("Statement 0")
	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 1.1")
	fmt.Println("Statement 2")
	fmt.Println("Statement 3")
	// SwitchStatement()
}