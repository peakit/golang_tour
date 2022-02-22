package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v *Vertex) Scale_PtrReceiver(f int) {
	// with pointer receiver
	v.x += f
	v.y += f
}

func (v Vertex) Scale_ValReceiver(f int) {
	// with value receiver
	v.x += f
	v.y += f
}

func Alternate_Scale(v *Vertex, f int) {
	v.x += f
	v.y += f
}

func main() {
	v := Vertex{1, 2}
	fmt.Printf("\nInitial value: v.x=%d, v.y=%d", v.x, v.y)

	v.Scale_PtrReceiver(10)
	fmt.Printf("\nScaled (with Ptr receiver) value: v.x=%d, v.y=%d", v.x, v.y)

	Alternate_Scale(&v, 10)
	fmt.Printf("\nScaled(alt) value: v.x=%d, v.y=%d", v.x, v.y)

	v.Scale_ValReceiver(10)
	fmt.Printf("\nScaled (with Value receiver) value: v.x=%d, v.y=%d", v.x, v.y)

	ptr := &v
	ptr.Scale_ValReceiver(10)
	fmt.Printf("\nScaled (with Value receiver) value: v.x=%d, v.y=%d", v.x, v.y)
}
