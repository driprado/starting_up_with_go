package main

import "fmt"

func main() {

	// Normal incremental loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Insertnig a breakpoint
	for n := 0; n < 10; n++ {
		if n > 5 {
			break // Breaks once n reaches 5
		}
		fmt.Println(n)
	}

	// Using continue will skiip an item
	a := 5
	for b := 0; b < 10; b++ {
		if b == a {
			continue // Will print all items skiiping 5 which is the value of b
		}
		fmt.Println(b)
	}

	// Other loop format
	alpha := 0
	for {
		fmt.Println(alpha)
		alpha = alpha + 1 // just another way to increment it
		if alpha == 10 {  // stops when it reaches 10
			break
		}
	}
	fmt.Println(alpha)

	// Looping through a string range:
	phrase := "tonga da mironga"
	for k, v := range phrase { // k=key(position in range), v=value(runes, need to be converted to string)
		fmt.Println(k, string(v))
	}
}
