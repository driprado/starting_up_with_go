package main

import "fmt"

func main() {
	// var i int = 10 -> longer way to declare
	// var i = 10 -> shorter way to declare (type infer)
	// i := 10 // Even shorter way t0 declare
	// fmt.Println(i)

	var i int //don't know the value yet. empty int variables in go assume value 0
	fmt.Println(i)
	i = 20 //new value
	fmt.Println(i)
	//Output:
	// 0
	// 20

}
