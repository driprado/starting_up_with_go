package main

import "fmt"

func main() {
	// true is a pre-declared constant in golang
	if true {
		fmt.Println("If there's no error or explicit conditions, is always TRUE")
	}
	if false {
		fmt.Println("that's FALSE")
	}
	if !true {
		fmt.Println("if is not TRUE than is FALSE")
	}
	if !false {
		fmt.Println("if not false than it's TRUE")
	}

	// initislise var and check if true in the same line:
	if x := 2; x == 2 {
		fmt.Println("x == 2, therefore, TRUE")
	}

	// Else statements
	if x := 2; x == 3 {
		fmt.Println("x == 2, therefore TRUE")
	} else {
		fmt.Println("x != 2, therefore FALSE")
	}

}
