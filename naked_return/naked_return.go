package main

import "fmt"

func split(sum int8) (x, y int8) {
	x = sum * 3 / 4
	y = sum - x
	return // naked return of named return parameters
}

func main() {
	fmt.Println(split(20))
}
