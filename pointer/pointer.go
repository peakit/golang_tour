package main

import "fmt"

func main() {
	i := 10
	var iptr *int
	iptr = &i
	fmt.Printf("Value of i is: %d\n", *iptr)
}
