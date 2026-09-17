Apparently Go has a function to create slices.
make([]string, length, capacity_optional)

below creates a slice with length five, all element are set to nil of the type, in below case its empty string ""
aSlice := make([]string, 5) 
["", "", "", "", ""]

if type if int
bSlice := make([]int, 5)

[0, 0, 0, 0, 0]


Slices of slcies

matrix := [][]int{
  []int{1, 2},
  []int{3, 4}
}

So slices can be created using two ways,

make or creating an array without specifying its capacity it becomes a slice

make([]T, lenght, capacity)

or 

s := []T{}
