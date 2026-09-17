package main

import "golang.org/x/tour/pic"

// Using make, having length at start as 0, but has capacity
func Pic(dx, dy int) [][]uint8 {
	picData := make([][]uint8, 0, dy)
	
	for i := 0; i < dy; i++ {
		picData = append(picData, make([]uint8, 0, dx))
	}
	
	for y, row := range picData {
		for x := 0; x < dx; x++ {
			row = append(row, uint8(x * y / 2))
		}
		picData[y] = row;
	}
	
	return picData
}

// Using slices with nil values, having length from start
func Pic(dx, dy int) [][]uint8 {
	picData := make([][]uint8, dy)
	
	// [nil, nil, nil]
	
	for y := range picData {
		row := make([]uint8, dx)
		
		for x := range row {
			row[x] = uint8(x * y / 2)
		}
		
		picData[y] = row;
	}
	
	return picData
}

func main() {
	pic.Show(Pic)
}
