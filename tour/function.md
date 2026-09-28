function are values in Go
meaning they can be passed around
like other values, can be returned or can be passed as arguments

func compute(fn func(int, int) float) int {
  return fn(2, 3)
}

compute(math.Pow)

myfunc := func(x, y int) int {
  return math.Sqrt(x*x + y*y)
}

lol

Go has closures like javascript

a function is bound to the variables it references from the outside.

e.g.:

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
