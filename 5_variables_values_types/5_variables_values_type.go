// You can use the function fmt.Println to print the type of a variable:
package main

import "fmt"

var n int = 50
var s string = "word"
var i int //don't know the value yet. empty int variables in go assume value 0

func main() {

	fmt.Println(i) // no value yet, 0
	i = 20         //new value
	fmt.Println(i) // 20

	fmt.Print(n)
	fmt.Printf("\t%T\n", n)
	fmt.Print(s)
	fmt.Printf("\t%T\n", s)
}
