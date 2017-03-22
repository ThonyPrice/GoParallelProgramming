/*
Task 2.1 - Loops and Functions, from https://tour.golang.org/flowcontrol/8
DD1396 - Parallel and concurrent programming @ KTH
Author: Thony Price
Last revision: 2017-03-22
*/

package main

import (
	"fmt"
	"math"
)

var value float64 = 2

// Sqrt1 computes square root of integer input using Newton's method 10 cycles
func Sqrt1(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}

// Sqrt2 computes square root of an integer input using Newton's method
// until value changes less than 0.000001 each iteration
func Sqrt2(x float64) float64 {
	it := 0
	z2 := x
	for z1 := 5.0; math.Abs(z2-z1) > 0.000001; {
		z1 = z2
		z2 = z1 - ((z1*z1 - x) / (2 * z1))
		it++
	}
	fmt.Printf("Delta method reqired %v iterations\n", it)
	return z2
}

func main() {
	fmt.Println("Math function	        :", math.Sqrt(value))
	fmt.Println("Newton's 10 iterations	:", Sqrt1(value))
	fmt.Println("Newton's stop@delta	:", Sqrt2(value))
}
