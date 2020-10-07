package main

import "fmt"

func main() {
	// #===#===#==#==#==# Slices #===#===#==#==#==# //

	// The make built-in function allocates and initializes an object of type slice(basically a list), map, or chan
	// The first argument is a type, not a value.
	mySlice := make([]string, 0)                      // Creates a slice of length 0
	fmt.Println("length of mySlice:", len(mySlice))   // Return 0
	mySlice = append(mySlice, "one")                  // Append "word" to slice
	fmt.Println("Content of position 0:", mySlice[0]) // Returns 1
	fmt.Println("length of mySlice:", len(mySlice))   // Returns 1
	mySlice = append(mySlice, "two")
	fmt.Println("length of mySlice:", len(mySlice)) // Returns 2
	fmt.Println("Contents mySlice:", mySlice)
	mySlice[0] = "zero" // Put "zero" in position 0 of slice
	fmt.Println("Contents mySlice:", mySlice)

	//Another Slice
	fmt.Println("============================")
	slice2 := make([]string, 2)                // Appendings for this list will go into position 2
	fmt.Println("contents of slice2:", slice2) // []Empty
	fmt.Println("============================")
	slice2 = append(slice2, "two") //Appended in position 2
	fmt.Println("contents of slice2:", slice2)
	fmt.Println("============================")
	fmt.Println("Contents of slice2[0]:", slice2[0])
	fmt.Println("Contents of slice2[1]:", slice2[1])
	fmt.Println("Contents of slice2[2]:", slice2[2]) //Returns zero
	fmt.Println("length of slice2:", len(slice2))
	fmt.Println("============================")
	fmt.Println("Adding 'zero' to slice2[0]")
	slice2[0] = "zero" // Inseting content into specific position in slice
	fmt.Println("Content of slice2[0]:", slice2[0])

	// Slices are not immutable!
	// Slices are reference types, they behave like pointers do; when slicing from a slice, the change in one slice affects the other
	sliceThree := []string{"x", "y", "z"}
	for k, v := range sliceThree { // For all key, values in sliceThree...
		fmt.Println(k, v)
	}
	//Output:
	// 0 x
	// 1 y
	// 2 z

	sliceFour := sliceThree[0:2] // Create sliceFour out of the elements 0-2 of sliceThree
	fmt.Println("Contents of sliceFour:", sliceFour)
	// Output:
	// Contents of sliceFour: [x y]

	sliceThree[0] = "a"                              // Changing stuff on sliceThree:
	fmt.Println("Contents of sliceFour:", sliceFour) // Will change consequentially sliceFour
	// Output:
	// Contents of sliceFour: [a y]

	var varSliceFive []string // Initializing an empty slice and attributing to a variable
	varSliceFive = sliceThree // Giving the variable the value of another slice
	fmt.Println("varSliceFive Contents:", varSliceFive)
	// Output:
	// varSliceFive Contents: [a y z]
	sliceThree[0] = "x"                                 // Changing stuff on sliceThree
	fmt.Println("varSliceFive Contents:", varSliceFive) // Consequently changes varSliceFive
	// Output: varSliceFive Contents: [a y z]

	// #===#===#==#==#==# Maps #===#===#==#==#==# //

	// Basically a dictionary
	// We also use make to create maps

	// Declaring an empty map
	mapOne := make(map[string]int) // Mapping a string to a integer
	mapOne["Marta"] = 17
	mapOne["Klose"] = 16
	mapOne["Ronaldo"] = 15
	fmt.Println("Contents of mapOne:", mapOne)              // Output: [Klose:16 Marta:17 Ronaldo:15]
	fmt.Println("Total goals for Marta:", mapOne["Marta"])  // Output: 17
	fmt.Println("Whats the value of xpto?", mapOne["xpto"]) // Key doesn't exist in map, returns 0
	// Testing if a key, value exist in a map:
	if value, ok := mapOne["Ronaldo"]; ok { // 'ok' is a boolean tha tests if that key/value pair exist in the map
		fmt.Println("Ronaldo exist and he score is:", value) // Output: Ronaldo exist and he score is: 15
	} else {
		fmt.Println("Sorry, player not found")
	}

	// Initializing a map with content:
	rank := map[string]int{ // Maps string to integer
		"Marta":   1,
		"Klose":   2,
		"Ronaldo": 3, // Add comma even in the last item
	}

	for key, value := range rank { // Print all key/value pair
		fmt.Println(key, value)
	}

	fmt.Println("Klose rank:", rank["Klose"]) // Output: Klose rank: 2
	delete(rank, "Klose")                     // Remove Klose entry
	fmt.Println("Klose rank:", rank["Klose"]) // Output: Klose rank: 0
	rank["Klose"] = 2                         // Add Klose back to map
	fmt.Println("Klose rank:", rank["Klose"]) // Output: Klose rank: 0
}
