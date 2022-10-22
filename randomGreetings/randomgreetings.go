package randomgreetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)


// return either of the formats specified inside the MESSAGE_OPTIONS
func randomFormat() string {
	// Slice(string slice) of message formats
	MESSAGE_OPTIONS:= []string {
		"Hi! %v welcome Bruh.",
		"Great To see ya again!, %v.",
		"%v, you're still here",
	}

	// rand.Intn(n), returns random number in the range [0, n), panics if n<0
	return MESSAGE_OPTIONS[rand.Intn( len(MESSAGE_OPTIONS) )]
}



func GetRandomGreeting(name string) (string, error) {
	// Return error if no name was given
	if(name == ""){
		return name, errors.New("name cannot be empty")
	}

	// create message using randomformat
	message:= fmt.Sprintf(randomFormat(), name)
	return message, nil
}

/*
	init(): special function along with main() which is used in the beginning of the program to initialize stuff
	initializing random with a seed ( constant seed ensure same sequence of random numbers each time )
	for a different sequence, initialize with seed: time.now().UnixNano(), 
*/
func init() {
	// rand.seed(25)
	rand.Seed(time.Now().UnixNano())
}
