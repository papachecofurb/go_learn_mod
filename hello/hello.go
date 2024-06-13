package main

import (
	"fmt"

	"github.com/papachecofurb/go_learn_mod/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Pedro")
	fmt.Println(message)
}
