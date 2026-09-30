What is interface in Golang?
It forces a type to have certain method to be implemented.

a type interface is defined as a set of method signatures.
IMPORTANT: A value of interface type can hold any value that implements those methods.

type FoolMethods interface {
  Random() int
  Shuffle() []int
}

func main() {
  var f FoolMethods
  // f can hold any value that implements Random() and Shuffle() methods

  f = &MyType{} // MyType must implement Random() and Shuffle() methods

}

Golang definition.
Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

Any type can implement methods of an interface, and it does not to explicitly declare that it implements that interface.
here is no explicit declaration of intent, no "implements" keyword.

So, basically, create a type that has methods that match the interface, and that type will implement the interface.
or kinda like work together, if a type has methods that match the interface, then that type implements the interface.

working example:

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
