struct is a collection of field.
Maybe it is same as in C.

declaration
type Vertex struct {
  X int
  Y int
}

usage

v := Vertex{1, 2}

accessing the fields

fmt.Println(v.X + v.Y)

pointers can also hold structs

pS := &v

fmt.Println(pS.X)


Stucts literals
Different ways to defining a struct

Below is a copy pasta from go.dev/tour

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}



