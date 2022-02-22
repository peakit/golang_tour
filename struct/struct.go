package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v1 := Vertex{1, 2}
	v2 := Vertex{3, 4}
	fmt.Println(v1, v2)
	fmt.Println(v1.X, v1.Y, v2.X, v2.Y)

	v3 := Vertex{X: 1}
	v4 := Vertex{}
	fmt.Println(v3.X, v3.Y, v4.X, v4.Y)
}
