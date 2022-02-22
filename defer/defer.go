package main

import "fmt"

func cleanUp1(i int) {
	fmt.Printf("i=%d in cleanUp function\n", i)
}

func cleanUp2(i int) {
	fmt.Printf("i=%d in cleanUp function2\n", i)
}

func main() {
	i := 90
	defer cleanUp1(i)
	i = i + 10
	defer cleanUp2(i) // deferred function calls are put on stack and executed in LIFO
}
