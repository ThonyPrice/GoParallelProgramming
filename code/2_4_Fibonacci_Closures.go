/*
Task 2.4 - Fibonacci Closures, from https://tour.golang.org/moretypes/26
DD1396 - Parallel and concurrent programming @ KTH
Author: Thony Price
Last revision: 2017-03-22
*/

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum1 := 0
	sum2 := 0
	sumn := 0
	return func() int {
		if sum1+sum2+sumn == 0 {
			sum1 = 1
			return sumn
		}
		sumn = sum1 + sum2
		sum1 = sum2
		sum2 = sumn
		return sumn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
