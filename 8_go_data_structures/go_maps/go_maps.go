//////////////////////////////////////// Readup:
// https://golang.org/ref/spec#Map_types - @done
// https://blog.golang.org/maps

//////////////////////////////////////// Notes:
// A map is an unordered group of elements.
// Maps are indexed by a set unique keys of one type that are mapped to values of another type.
// MapType = "map" "[" KeyType "]" ElementType .
// Ex: map[string]int // maps strings to integers
// A new, empty map value is made using the built-in function make, which takes the map type and an optional capacity hint as arguments:
// Ex: make(map[string]int, 100)
// The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, with the exception of nil maps.
// A nil map is equivalent to an empty map except that no elements may be added.

package main

import "fmt"

func main() {

	//////////////////////////////////////// Example 1:
	// phoneList maps names to phone numbers
	phoneList := map[string]int{
		"Anna":     404765123,
		"Beatrice": 401234321,
		"Carol":    404453655,
		"Danielle": 404656987,
		"Ellie":    404434982,
	}

	fmt.Println(phoneList)          // map[Anna:404765123 Beatrice:401234321 Carol:404453655 Danielle:404656987 Ellie:404434982]
	fmt.Println(phoneList["Ellie"]) // 404434982

	// If you query for a key that does not exist, if values are integers, map will return 0 instead of an error:
	fmt.Println(phoneList["Fiona"]) // 0

	// Verifying if a key/value pair exist in the map:
	// To check if a key/value pair exist in the map, you can use the boolian `ok`:
	value, ok := phoneList["Fiona"]
	fmt.Println(value) // 0 // values is 0 as the key "Fiona" doesn't exist
	fmt.Println(ok)    // false // value is not ok, it does not exist, nor the key

	// Verifying an existing key:
	value, ok = phoneList["Anna"]
	fmt.Println(value) // 404765123
	fmt.Println(ok)    // true

	// Whe can than use an if statement to actionable code only if 'ok' == 'true':
	if value, ok = phoneList["Fiona"]; ok {
		fmt.Println("The phone number is:", value)
	} else {
		fmt.Println("Invalid key") // in this case this will be printed
	}

	//////////////////////////////////////// Example 2:
	// Use the builtin 'make' function to initialise an empty map:
	// The make function allocates and initializes a hash map data structure and returns a map value that points to it
	recordGoals := make(map[string]int)

	recordGoals = map[string]int{
		"Pele":    77,
		"Neymar":  64,
		"Ronaldo": 62,
	}

	// Individually adding key/vaues to a map:
	recordGoals["Marta"] = 108
	fmt.Println(recordGoals)        // map[Marta:108 Neymar:64 Pele:77 Ronaldo:62]
	theBest := recordGoals["Marta"] // Assign the value of Marta to theBest variable (outside map)
	fmt.Println(recordGoals)        // map[Marta:108 Neymar:64 Pele:77 Ronaldo:62] // not alrered
	fmt.Println(theBest)            //108

	// len function returns the number of items in a map:
	fmt.Println(len(recordGoals)) // 4

	// delete function removes an entry from the map:
	delete(recordGoals, "Neymar")
	fmt.Println(recordGoals) // map[Marta:108 Pele:77 Ronaldo:62]

	//////////////////////////////////////// Example 3:
	// You're a headhunter and have a map of programmers and their preferred programming languages
	// You want to know which are the programmers that know Golang:

	// @todo NEED DECENT VAR NAMES, RENAME STUFF!!!
	type Coders struct {
		Name      string
		Languages []string
	}

	coder := []*Coders{
		{
			Name:      "Ada",
			Languages: []string{"math", "differential", "calculus", "discrete"},
		},
		{
			Name:      "Hoper",
			Languages: []string{"assembly", "cobol", "contran"},
		},
		{
			Name:      "Borg",
			Languages: []string{"assembly", "cobol", "C", "bash"},
		},
		{
			Name:      "Hamilton",
			Languages: []string{"assembly", "corona"},
		},
		{
			Name:      "Korbes",
			Languages: []string{"ruby", "python", "go"},
		},
		{
			Name:      "Prado",
			Languages: []string{"ruby", "python", "go"},
		},
	}

	fluent := make(map[string][]*Coders)
	for _, programer := range coder {
		for _, lan := range programer.Languages {
			fluent[lan] = append(fluent[lan], programer)
		}
	}

	for _, c := range fluent["python"] {
		fmt.Println(c.Name, "codes in python")
	}
}
