package main

// func divAndMod(a int, b int) (int, int, error) { // Tekes two int, returns two ints and one error
// 	if b == 0 {
// 		return 0, 0, errors.New("Can't devide by zero") // If b = 0 returns error
// 	}
// 	return a / b, a % b, nil // Else return results of operation
// }

func main() {

	// 	if len(os.Args) < 3 { // If not all parameters are passed, exit 1
	// 		fmt.Println("Expected two input parameters")
	// 		os.Exit(1)
	// 	}

	// 	a, err := strconv.ParseInt(os.Args[1], 10, 64) // Test if argument a can be parsed into int correctly
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	b, err := strconv.ParseInt(os.Args[2], 10, 64) // Test if argument b can be parsed into int correctly
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	div, mod, err := divAndMod(int(a), int(b)) // Calling function divAndMod
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Println("%d / %d == %d and %d %% %d == %d\n", a, b, div, a, b, mod)

	// Custom message errors:
	type MyError struct { // Any variable or type that starts with an uppercase letter is exported from that package
		A       int
		B       int
		Message string
	}

	// func (err *MyError) Error() string // Can't can't understand this shit right now

}
