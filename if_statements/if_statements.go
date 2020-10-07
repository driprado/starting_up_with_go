package main

import "fmt"

func main() {
	// Simple if statement
	n := 5
	if n > 5 {
		fmt.Println("n is bigger than 5")
	} else {
		fmt.Println("n is less than or equal to 5")
	}
	// Output
	// n is less than or equal to 5

	//If and statement
	a := 20
	if b := a * 2; b > 20 { // If b is equal a*2 AND b >20
		fmt.Println("b is bigger than 20\tThis is b:", b, "This is a:", a)
	} else {
		fmt.Println("b is less than or equal 20")
	}
	//fmt.Println(b) // Variables defined inside if statements only exists within the statements, this line of code would generate an error.

	// For statement

}
