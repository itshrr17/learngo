# Zero Values
variables declared without explicit value, are initialized to zero value
each type as its own zero value
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

# Type conversion
type conversion are done using this expression Type(value)

e.g.: float64(i)

var i int = 42
var f float64 = float64(i)

or 

i := 42
f := float64(i)