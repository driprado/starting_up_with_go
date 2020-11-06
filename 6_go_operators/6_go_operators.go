package main

import (
	"fmt"
)

func main() {
	//// Arithmetic operators
	// + sum
	x := 2
	y := 5
	z := x + y
	fmt.Println(z)

	// - difference
	z = x - y
	fmt.Println(z)

	// * product
	z = x * y
	fmt.Println(z)

	// / quotient
	z = x / y
	fmt.Println(z)

	// % remainder
	z = x % y
	fmt.Println(z)

	//// Comparison operators
	// ==    equal
	if x+y == 7 {
		fmt.Println("true statement")
	}

	// !=    not equal
	if x-y != -3 {
		fmt.Println("this will never print")
	}

	// <     less
	if x < y {
		fmt.Println("also true statement")
	}

	// <=    less or equal
	if x <= y {
		fmt.Println("still true statement")
	}
	// >     greater
	if x > y {
		fmt.Println("won't print")
	}
	// >=    greater or equal
	if x > y {
		fmt.Println("won't print")
	}

	//// Logical Operators
	// &&    conditional AND
	fmt.Println("true && true\t", true && true)   // true
	fmt.Println("true && false\t", true && false) // false

	// ||    conditional OR
	fmt.Println("true || true\t", true || true)   // true
	fmt.Println("true || false\t", true || false) // true

	// !  NOT
	fmt.Println("!true", !true)   // false
	fmt.Println("!false", !false) // true

}
