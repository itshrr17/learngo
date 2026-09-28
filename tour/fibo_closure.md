// I cant believe i struggled 10 mins, figuring out how to do this


package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	curr := 0
	next := 1
	
	return func() int {
		t := curr
		curr = next
		next = t + next
		return t
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
