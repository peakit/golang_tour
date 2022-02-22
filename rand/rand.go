package main

import (
	"fmt"
	"math"
	"math/rand"
)

func sum(a int8, b int8) int8 {
	return a + b
}

func new_sum(a, b int8) int8 {
	return a + b
}

func main() {
	fmt.Printf("This %d is a random value!\n", rand.Intn(10))
	fmt.Printf("Pi: %f is a floating point number\n", math.Pi)

	x := sum(10, 20)
	fmt.Printf("Sum: %d\n", x)
	fmt.Println(new_sum(35, 20))
}
