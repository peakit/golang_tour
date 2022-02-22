package main

import "fmt"

func main() {
	arr := []int{
		1, 2, 3, 4,
	}
	for idx, val := range arr {
		fmt.Printf("At index=%d, the value is %d\n", idx, val)
	}

	for _, val := range arr {
		fmt.Printf("Value:%d\n", val)
	}
}
