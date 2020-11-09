// https://blog.golang.org/slices-intro
// https://blog.golang.org/slices
// https://tour.golang.org/moretypes/7

package main

import "fmt"

func main() {
	//// Declare slice with predetermided size and give valaues by index position
	var s [3]int
	s[0] = 10 // Add value to item in position 0
	s[1] = 11 // Add value to item in position 1
	s[2] = 12 // Add value to item in position 2

	fmt.Println(s)                // print slice
	fmt.Println(s[0], s[1], s[2]) // print items of slice one by one
	for index, value := range s { // print iterating through slice
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	//// Initialise/Declare sclice with values
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range sl { // print iterating through slice
		fmt.Println(v)
	}

	//// Slicing slices
	fmt.Println(sl[0])   // prints what's in position 0
	fmt.Println(sl[0:])  // prints what's in position 0 to the end of slice
	fmt.Println(sl[1:5]) // prints from position 1 to/but not included 4

	//// Slice slice into other slice
	five := sl[0:5]
	fmt.Println(five)

	//// Appending items to slices
	// func append(slice []Type, elems ...Type) []Type // Takes a slice of any type, append a variatic amount of elements of any type, returns a slice
	sl = append(sl, 10, 11, 12)
	fmt.Println(sl)

	//// Deleting items from slices
	// To remove an item from slice we have to "cut it out", as in using append and "jumping" the item you need to delete
	sl = append(sl[0:10], sl[11:]...) // "jump item 10, therefore removing it"
	fmt.Println(sl)
	// If you have to do this multiple times, you can create a delete function: "deleteFromSlice"
	sl = deleteFromSlice(sl, 2)

	//// Using make to create a slice
	// https://golang.org/pkg/builtin/#make
	// make alows you to set lenght and capacity of slice
	// Slice: The size specifies the length. The capacity of the slice is
	// equal to its length. A second integer argument may be provided to
	// specify a different capacity
	// make([]int, 0, 10) allocates an underlying array of size 10 and returns a slice of length 0 and capacity 10
	x := make([]int, 10, 100)
	fmt.Println(x)      // [0 0 0 0 0 0 0 0 0 0] make initializes the slice with value 0 if integer
	fmt.Println(len(x)) // 10
	fmt.Println(cap(x)) // 100

	//// Multidimentional slices
	// Multidimentional slices are slice of slices
	sample := [2][3]int{{1001, 1002, 1003}, {1004, 1005, 1006}}

	fmt.Printf("Number of rows in array: %d\n", len(sample))
	fmt.Printf("Number of columns in array: %d\n", len(sample[0]))
	fmt.Printf("Total number of elements in array: %d\n", len(sample)*len(sample[0]))

	for _, row := range sample { // for each slice in sample
		for _, item := range row { // for each item inside of each slice in sample
			fmt.Println(item)
		}
	}

}

func deleteFromSlice(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
