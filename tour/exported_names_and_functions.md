## Exported Names

a name is exported when it starts with capital letter
e.g.: Hello, Pi, Rand

If the first letter is not capital they will not be exported from the package.
e.g.: hello, pi, rand

So it is important, when exported capitalize first letter, if dont want export let it start with small letter.

## Functions
Function can take zero or more arguments.
e.g.:
arguments types come after the arguments.

func add(x int, y int) int {
  return x + y
}

func Hello() {
  fmt.Println("Hello")
}

we can also omit the type of the prev argument if it is the same as the next one.
e.g.:
func add(x, y int) int {
  return x + y
}

## Multiple Return Values
Go functions can return multiple values.

e.g.:
func swap(x, y string) (string, string) {
  return y, x
}

## Named Return Values
Go also supports named return values.

e.g.:
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

see in return statement we can omit the return values because they are already named.
it is called naked return, but it is not recommended to use naked return in long functions because it can be confusing.

