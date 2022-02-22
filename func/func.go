package main

import (
	"fmt"
	"math"
)

/**
Func which takes another func as argument. We are just abstracting the func but cannot abstract the operating
*/
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func fibonacci() func() int {
	prev := 0
	next := 1
	old_next := 0
	return func() int {
		old_next = next
		next = next + prev
		prev = next
		return next
	}
}

func main() {
	fmt.Println(compute(math.Pow))
	// declaring a func to be passed in as argument
	hypotenuse := func(x float64, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypotenuse))

	// fibonacci series using closures
	fmt.Println("\n\nFibonacci series")
	fibo := fibonacci()
	for i := 0; i < 5; i++ {
		fmt.Println(fibo())
	}
}
