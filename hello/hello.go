package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Hemraj")
	fmt.Println(message)
}
