package main

import "fmt"

// Just printing a sum
func main() {
	addNumbers(2, 2)
	addNumbers(7, 3)
}
func addNumbers(x int, y int) {
	fmt.Println(x + y)
}

// // Using return
// func main() {
// 	result := addNumbers(2, 8) // := is used only on the first time you declare the variable
// 	fmt.Println(result)

// 	result = addNumbers(4, 6) // The remaining times the variable appear, = is enough
// 	fmt.Println(result)

// 	result = addNumbers(7, 3)
// 	fmt.Println(result)
// }

// func addNumbers(x int, y int) int {
// 	return x + y
// }

// // Passing and Returning multiple values fom a function
// func multiplyAndDevide(x int, y int) (int, int) { //If you don't declare all arguments the function takes you get an error
// 	return x * y, x / y
// }

// func main() {
// 	mult, div := multiplyAndDevide(10, 2) // Each result is associated with a variable
// 	fmt.Println(mult, div)

// 	mult, _ = multiplyAndDevide(4, 2) // Can't declare a variable that is not used, so if we want to print only one result
// 	fmt.Println(mult)                 // we can use _ to ignore the other variable

// 	_, div = multiplyAndDevide(5, 5)
// 	fmt.Println(div)
// }

// // Simple function that add 1 to any given int number
// func plusOne(a int) int {
// 	return a + 1
// }

// func main() {
// 	fmt.Println(plusOne(1))
// }

// // Passing the function and value for a variable
// func plusOne(a int) int {
// 	return a + 1
// }

// func main() {
// 	var_func := plusOne
// 	fmt.Println(var_func(5)) // Here we use the variable instead of the function name
// }

// //Creating a function inside another function
// func main() {
// 	subFunc := func(a int) int { // A function inside another has to be declared as an annonimous function assigned to a variable.
// 		return a + 1 // The function declared here is not accessible outside the main function
// 	}
// 	fmt.Println(subFunc(4))
// }
