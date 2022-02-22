package main

import "fmt"

type Runnable interface {
	Run()
}

type Animal struct {
	name string
}

// this means Animal struct implements the
// interface Runnable as it puts a definition
// of run() with a value receiver
func (a Animal) Run() {
	fmt.Printf("%s is running!\n", a.name)
}

func main() {
	dog := Animal{name: "Sultan"}
	dog.Run()

	var dogI Runnable
	dogI = &dog
	fmt.Printf("Let's try again using %v of type %T\n", dogI, dogI)
	dogI.Run()

	_, ok := dogI.(Animal)
	if ok {
		fmt.Println("dogI is holding concrete type of Animal")
	} else {
		fmt.Println("dogI is not holding concrete type of Animal")
	}

	tVal1 := dogI.(*Animal)
	fmt.Println("tVal: ", tVal1)
}
