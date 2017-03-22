/*
Task 2.2 - Slices, from https://tour.golang.org/moretypes/18
DD1396 - Parallel and concurrent programming @ KTH
Author: Thony Price
Last revision: 2017-03-22
*/

package main

import "golang.org/x/tour/pic"

// Pic returns a 2-dimensional array of uint8 values
func Pic(dx, dy int) [][]uint8 {
	dY := make([][]uint8, dy)
	for i := range dY {
		dY[i] = make([]uint8, dx)
		for j := range dY[i] {
			dY[i][j] = uint8((i * j))
		}
	}
	return dY
}

func main() {
	pic.Show(Pic)
}
