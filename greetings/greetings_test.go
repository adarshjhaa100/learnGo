package greetings  // same package, can invoke the function defined in other files of the same package without import


import (
	"testing"
	"regexp"
	"fmt"
)

func TestGreeting(t *testing.T) {
	name:= "Human"
	want:= regexp.MustCompile(`\b`+name+`\b`) // with message matching %name%

	fmt.Println("Want regexp:")
	fmt.Println(want) 

	msg, err:= GreetMe(name)
	// msg = "32324"
	if !want.MatchString(msg) || err!=nil {
		// backquote string is like multiline string in python and ignore special characters ( used for putting in regexps )

		t.Fatalf(`GreetMe(%v)=%v, %q, want match for %#q, nil`, msg, msg, err, want)
	} 

}


// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestGreetingloEmpty(t *testing.T) {
    msg, err := GreetMe("")
    if msg != "" || err == nil {
        t.Fatalf(`GreetMe("") = %q, %v, want "", error`, msg, err)
    }
}


