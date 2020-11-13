// Example 2: Value receiver
package main

import (
	"fmt"
)

type operation struct {
	X, Y int
}

// (o operation) ==> receiver, a reference to the struct operation, bounds the method to the operation type
// myProduct ==> method name
// int ==> return type
func (o operation) myProduct() int {
	return ((o.X * o.X) * (o.Y * o.Y)) // doesn't change struct, only makes calculations on top of its values
}

func main() {

	o := operation{2, 5}       // passing values to operation struct
	fmt.Println(o.myProduct()) // applying 'myProduct' method to 'o'

}
