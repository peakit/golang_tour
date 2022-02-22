package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// an equivalent while loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// an equivalent infinite loop
	for {
		if sum < 5000 {
			sum += sum
		}
		break
	}
	fmt.Println(sum)

}
