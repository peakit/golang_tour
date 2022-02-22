package main

import "fmt"

var some_var1, some_var2 bool

// similar to imports, vars have a "factored" block
var (
	Z bool = false
)

// similarly we have factored block for const as well
const (
	NUM_99     = 99
	TRUTH_BOOL = true
)

func main() {
	var some_var3 int
	fmt.Println(some_var1, some_var2, some_var3, Z)

	// with initialization
	var some_var4, some_var5 bool = true, false
	fmt.Println(some_var4, some_var5)

	// with initialization but type is omitted now
	var some_var6, some_var7 = 1, false
	fmt.Println(some_var6, some_var7)

	// short hand for variable "declaration-and-initialization"
	some_var8 := true
	fmt.Println(some_var8)

	// const MY_PI := 3.14 (const cannot be declared using short hand)
	const MY_PI = 3.14
	fmt.Println(MY_PI)

}
