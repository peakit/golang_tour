package main

import "fmt"

func fibo() func() int {
	first := 0
	second := 1

	return func() int {
		next := first + second
		first = second
		second = next

		return next
	}
}

func main() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
