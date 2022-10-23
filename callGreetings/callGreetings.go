package main

import (
	"fmt"
	"learningo/greetings"
	"log"
	"learningo/randomgreetings"
)

func handleAndGreet(name string) {
	message, error:= greetings.GreetMe(name)
	fmt.Println(message)
	// message, error = randomgreetings.GetRandomGreeting(name)
	
	listName:= []string {
		"Pikachu",
		"Bulbasaur",
		"Squirtle", //dont forget this last comma
	}

	messages,error := randomgreetings.GetMultipleRandomGreetings(listName)

	// if error is not nil, something went wrong
	if(error!=nil){
		// handle error message
		log.Fatal(error)
	}
	
	fmt.Println(messages) // can directly print a map to console

}


func main() {
	name1:= "Bulbasaur"
	// name2:= ""
	
	/*
		Set prefix of logger and by default time will be printed
	*/
	log.SetPrefix("callGreetings: ")


	// call function from other package
	handleAndGreet(name1)
	// handleAndGreet(name2)

	// fmt.Println("Hello World")
}
