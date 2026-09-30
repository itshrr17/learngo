Go does not have classes.
But we can define methods on types.
It kinda reminds me of prototypes in JavaScript.

type struct Vertex {
  X, Y float64
}

func (v Vertex) Neg() float64 {
  v.X = v.X * -1
  v.Y = v.Y * -1
  return v
}

func (v Vertex) Neg() Vertex {
  v.X = v.X * -1
  v.Y = v.Y * -1
  return v
}

usage 
v := Vertex{1.3, 3.5}
v = v.Abs()
v = v.Neg()

Explanation
In  func (v Vertex)
v Vertex is a receiver, a special argument
This way go knows this method belongs to Vertex type


Method is just a function with a receiver argument

Difference

func (v Vertex) Neg() Vertex {
  v.X = v.X * -1
  v.Y = v.Y * -1
  return v
}

func Neg(v Vertex) Vertex {
  v.X = v.X * -1
  v.Y = v.Y * -1
  return v
}

Methods can be declared on non struct types.

type localFloat float64 // this must be in the same package to work

func (f localFloat) Double() localFloat {
	return f * 2
}

xyc := localFloat(3.14)
fmt.Println(xyc.Double())

But cannot be declared on local types like, func (f float64); Wrong;

Pointer receiver or Pointers and functions

func (v* Vertex) Neg() float64 {
  v.X = v.X * -1
  v.Y = v.Y * -1
}

v.Abs() // changes that pointer, because it gets pointer not the value

or

func Abs(v* Vertex) float64 {
  v.X = v.X * -1
  v.Y = v.Y * -1
}

Abs(&v)

When to choose pointer receiver, when we want to modify the value 
or avoid copying the value, can be efficient when dealing with large structs