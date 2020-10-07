package main

import "fmt"

func main() {
	// var s string // Declared without value, assumes string list of lenght 0
	// s = "Hello world!"
	// fmt.Println(s)
	// Output:
	// Hello world!

	s := "Hello world!" //Shorter declaration
	fmt.Println(s)
	// Output:
	// Hello world!

	//Escaping special characters:
	s_two := "This is a backslash \\"
	fmt.Println(s_two)

	s_three := `A backthick also escape characters \, and also emojis! 🥓`
	fmt.Println(s_three)

	//New lines and tabs
	s_four := "To be or not to be\nThat is the question\t!!!"
	fmt.Println(s_four)
	// Output
	// To be or not to be
	// That is the question    !!!

	// Using plus operator
	s5 := "It's time for o cup of"
	s6 := "☕️"
	s7 := s5 + " " + s6
	fmt.Println(s7)

	//Selecting substrings out of strings
	tobe := "to be or not to be that is the question"
	t1 := tobe[0:5]
	t2 := tobe[13:18]
	fmt.Println(t1, t2)
	// Output:
	// to be to be

}
