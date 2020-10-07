package main

import "fmt"

func main() {
	var i int8 = 20
	var f float32 = 5.6
	// fmt.Println(i + f) -> invalid operation: i + f (mismatched types int8 and float32)
	fmt.Println(float32(i) + f) //Convert i type to fix
	// Output:
	// 25.6
	fmt.Println(i + int8(f)) // Converting the other variable
	// Output:
	//25
}
