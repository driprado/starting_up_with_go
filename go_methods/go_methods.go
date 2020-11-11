//// Reading
// https://www.youtube.com/watch?v=i3o4G4bmqPc 			@todo: create example
// https://www.youtube.com/watch?v=93f9_bJQdHk			@todo: create example@todo-make example
// https://golang.org/ref/spec#Method_declarations	@done
// https://golang.org/ref/spec#Method_sets
// https://golang.org/doc/faq#methods_on_values_or_pointers

//// Notes:
// ToSelf: Read code below then comeback to Notes
// A method is a function with a receiver.
// A method declaration associates the method with the receiver (ex: struct)

// There are two types of methods in Go:
// - Pointer receivers
// 		.	capable of changing values in the receiver's base type (ex: structs)
// - Value receivers
// 		. perform operations with receiver values, but doesn't permanently change the receiver's base type

package main

import (
	"fmt"
)

// Example 1: // @todo: change types and variables names, make it meaningful
type myStruct struct {
	X, Y int
}

// (s myStruct) ==> receiver, a reference to the struct above
// myOperation ==> method name
// float64 ==> return type
func (s myStruct) myOperation() int {
	return ((s.X * s.X) * (s.Y * s.Y))
}

func main() {
	s := myStruct{2, 5}
	fmt.Println(s.myOperation())
}
