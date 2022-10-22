package greetings

import (
	"fmt"
	"errors"
)

// function declaration, should start with an uppercase letter in case want to import
// return a tuple of string value and error type
func GreetMe(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty string found")
	}

	/* sprintf formats:
	v:	default format, almost same as C, didn't find anything different so far
	*/
	message := fmt.Sprintf("Hello %v, welcome to this world!", name)

	// return nil as a placeholder if there's no error
	return message, nil
}
