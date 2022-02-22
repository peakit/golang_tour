package main

import "fmt"

func main() {
	var a [5]string
	a[0] = "hello"
	a[1] = "world"
	a[2] = "to"
	a[3] = "newbies"
	a[4] = "of Golang"

	fmt.Println(a)
	fmt.Println(a[0:2])

	theBollywoodStars := []string{
		"Salman",
		"Shahrukh",
		"Aamir",
	}
	fmt.Println(theBollywoodStars)
	topStars := theBollywoodStars[1:]
	// after appending to a slice
	topStars = append(topStars, "Kajol")
	fmt.Println(topStars, theBollywoodStars)

	theVertices := []struct {
		x int
		y int
	}{
		{1, 2},
		{2, 3},
		{3, 4},
		{5, 6},
	}
	fmt.Println(theVertices)
	// with slicing
	fmt.Println(theVertices[:1], theVertices[2:])

	// length and capacity of slices
	firstTwo := theVertices[:2]
	fmt.Printf("len:%d and capacity:%d for slice:%v\n", len(firstTwo), cap(firstTwo), firstTwo)
	// since len is less than capacity, the slice could be extended
	firstThree := theVertices[:3]
	fmt.Printf("len:%d and capacity:%d for slice:%v\n", len(firstThree), cap(firstThree), firstThree)

	// dynamic arrays by slicing using 'make' function
	dynaArray := make([]int, 10)
	fmt.Println(dynaArray)
	fmt.Println(dynaArray[:3], len(dynaArray[:3]), cap(dynaArray[:3]))

	// slices of slices
	_2dArray := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	fmt.Println(_2dArray)

	for idx, val := range _2dArray {
		fmt.Printf("_2dArray[%d]:%v\n", idx, val)
	}

	for idx, _ := range _2dArray {
		fmt.Println(idx)
	}
}
