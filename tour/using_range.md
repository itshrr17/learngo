Range, looping over an array, slice or a map

for index, value := range array/slice/map {
  ...
}

var pow = [8]int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

  // when we dont need index or value, we replace it with an underscore "_"
  for _, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
  }

  for i := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
  }
}


