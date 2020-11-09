// https://golang.org/ref/spec#Map_types
// https://blog.golang.org/maps
// Unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type.
// map[KeyType]ValueType Ex: map[string]int

package main

import "fmt"

func main() {
	// Example 1
	// phoneList maps names to phone numbers
	phoneList := map[string]int{
		"Anna":     404765123,
		"Beatrice": 401234321,
		"Carol":    404453655,
		"Danielle": 404656987,
		"Ellie":    404434982,
	}
	fmt.Println(phoneList)          // Print all Keys and values
	fmt.Println(phoneList["Ellie"]) // Print value for key "Ellie"

	// If you query for a key that does not exist, in the case of values being inegers, map will return 0 instead of an error:
	fmt.Println(phoneList["Fiona"]) // 0
	// To check if the key/value you're querying for exist in the map, you can use the boolian `ok`:
	value, ok := phoneList["Fiona"]
	fmt.Println(value) // 0
	fmt.Println(ok)    // false

	value, ok = phoneList["Anna"]
	fmt.Println(value) // 404765123
	fmt.Println(ok)    // true

	// Whe can than use an if statement to actionable code only if 'ok' == 'true':
	if value, ok = phoneList["Fiona"]; ok {
		fmt.Println("The phone number is:", value)
	} else {
		fmt.Println("Invalid key")
	}

	// Example 2
	// Use the builtin make function to initialise an empty map:
	// The make function allocates and initializes a hash map data structure and returns a map value that points to it.
	recordGoals := make(map[string]int)

	recordGoals = map[string]int{
		"Pele":    77,
		"Neymar":  64,
		"Ronaldo": 62,
	}

	// Individually adding key/vaues to a map:
	recordGoals["Marta"] = 108
	fmt.Println(recordGoals)
	theBest := recordGoals["Marta"] // Assign the value of Marta to theBest variable (outside map)
	fmt.Println(recordGoals)
	fmt.Println(theBest) //108

	// len function returns on the number of items in a map:
	fmt.Println(len(recordGoals)) // 4

	// delete function removes an entry from the map:
	delete(recordGoals, "Neymar")
	fmt.Println(recordGoals) // map[Marta:108 Pele:77 Ronaldo:62]

	//// Example 3 - Exploiting zero values
	// A gym has as promotion to give 2 weeks free for new members only.
	// We need to make check if an email has already been signed

	type Member struct {
		Next  *Member
		Email string
	}

	var member *Member

	registered := make(map[*Member]bool)
	for person := member; person != nil; person = person.Next {
		if registered[person] {
			fmt.Println("already registered")
			break
		}
		registered[person] = true
		fmt.Println(person.Email)
	}

	// Preencher Member com varias pessoas aqui!!!!
	// Depois continuar daqui: https://blog.golang.org/maps
}
