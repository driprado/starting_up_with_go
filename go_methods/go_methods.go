package main

import (
	"fmt"
	"math"
)

type myStruct struct {
	X, Y float64
}

// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.

func (s myStruct) myOperation() float64 { // Receiver: (s myStruct), // Method: myOperation(), //Returns: float64
	return math.Sqrt(s.X*s.X + s.Y*s.Y)
}

func main() {
	s := myStruct{3, 4}
	fmt.Println(s.myOperation())
}
