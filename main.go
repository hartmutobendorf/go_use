package main

import (
	"fmt"

	"github.com/hartmutobendorf/go_mod"
)

func main() {
	// Get a greeting message and print it.
	message := go_mod.Hello("Gladys")
	fmt.Println(message)
}
