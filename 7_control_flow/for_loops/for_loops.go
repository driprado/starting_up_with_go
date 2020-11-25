package main

import "fmt"

func main() {

	// Golang for loop skeleton:
	// for init; condition; post {}
	// ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
	// Condition = Expression .

	// Normal incremental loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// break
	// Insertnig a breakpoint
	for n := 0; n < 10; n++ {
		if n > 8 {
			break // Breaks once n reaches 5
		}
		fmt.Println(n)
	}

	//contirue
	// Using continue will skip an item
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

	// Nested Loops
	for o := 0; o <= 10; o++ {
		fmt.Println("OUT")
		for i := 0; i <= 5; i++ {
			fmt.Println("in")
		}
	}

	// For with implicit condition
	n := 0
	for { //forever?
		if n > 20 {
			break
		}
		fmt.Println(n)
		n++
	}
	fmt.Println("DONE")
}
