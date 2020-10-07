package myfruits

import "strings"

//  Go's testing package

// testing.T has methods that allow you to inform test runs of errors or failures
// t.Fail() ==> This thest failed but keep running tests
// t.Errorf ==> Log the error than mask test as failed

var myFruits = []string{"tucuma", "dao", "jambo"}

// IsMyFruits verifies if a fruit is part of myFrutis slice
func IsMyFruits(fruit string) bool { // Takes a string named fruit, returns a bool
	for _, value := range myFruits {
		if strings.Contains(fruit, value) {
			return true
		}
	}
	return false
}
