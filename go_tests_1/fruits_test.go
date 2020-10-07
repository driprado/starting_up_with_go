package myfruits

import "testing"

// TestFruits pass a string and tests if is contained in myFruits
func TestFruits(t *testing.T) {
	fruit := "tucuma"
	if !IsMyFruits(fruit) {
		t.Fail()
	}
}
