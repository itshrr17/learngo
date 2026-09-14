Defer
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}


so delay the function execution but evaluate arguments immediately.
a fun way to use.

no need to declare variables for later use i think.

defer stacks LIFO
defer 1
defer 2 
defer 3

output
3
2
1

works like stack