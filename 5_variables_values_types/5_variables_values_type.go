// // 1 - Declaring variables and types
// // Go is a static programming language, meaning if you declare a variable with a specific type, that type doesn't change
// // You can use the function fmt.Println to print the type of a variable:
// package main

// import "fmt"

// var n int = 50
// var s string = "word"
// var i int //don't know the value yet. empty int variables in go assume value 0

// func main() {

// 	fmt.Println(i) // no value yet, 0
// 	i = 20         //new value
// 	fmt.Println(i) // 20

// 	fmt.Print(n)
// 	fmt.Printf("\t%T\n", n) // print type of variable n
// 	fmt.Print(s)
// 	fmt.Printf("\t%T\n", s) // print type of variabel s

// 	// You can also print number variables into other number systems:
// 	fmt.Printf("\n%b", n)             // print n in binary
// 	fmt.Printf("\n%x", n)             // print n in hexadecimal
// 	fmt.Printf("\n%o", n)             // print n in octal
// 	fmt.Printf("%o\t%b\t%x", n, n, n) // print all in same line
// }

// // 2 - Creating your own type
// package main

// import "fmt"

// type fruit string // define type fruit
// var a fruit       // declare var a of type fruit

// func main() {
// 	a = "apple"
// 	fmt.Print(a, "has type")
// 	fmt.Printf("\t%T", a)
// }

// 3 - Converting types
package main

import "fmt"

type fruit string // define type fruit
var a fruit       // declare var a of type fruit

func main() {
	a = "apple"
	fmt.Print(a, "has type")
	fmt.Printf("\t%T\n", a)
	b := string(a) // Converts a into string which is the underlying type of fruit
	fmt.Print(a, "has type")
	fmt.Printf("\t%T", b)
}
