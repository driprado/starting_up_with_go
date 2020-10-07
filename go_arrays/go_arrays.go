package main

import "fmt"

func main() {
	var myarray [3]int // Declare array of integers and lenght 3
	myarray[0] = 10    // Add value to item in position 0
	myarray[1] = 11    // Add value to item in position 1
	myarray[2] = 12    // Add value to item in position 2

	fmt.Println(myarray, myarray[0], myarray[1], myarray[2])
	// Output
	// [10 11 12] 10 11 12
}
