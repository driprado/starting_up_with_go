// https://golang.org/ref/spec#Switch_statements
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// A switch statement is a shorter way to write a sequence of if - else statements.
	// https://tour.golang.org/flowcontrol/9
	// In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped. If no case matches and there is a "default" case, its statements are executed. There can be at most one default case and it may appear anywhere in the "switch" statement. A missing switch expression is equivalent to the boolean value true.
	// If no case matches and there is a "default"
	// "fallthrough" statement indicate that control should flow from the end of this clause to the first statement of the next clause.

	// Switch OS
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch on variable value
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Switch on logic condition
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // match two conditions
		fmt.Println("It's the weekend")
	default: // default sounds like an else here
		fmt.Println("It's a weekday")
	}

	// Switch doesn't need to have a name
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Type switch
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// Switch same as if has true as implicit constant
	switch {
	case false:
		fmt.Println("FALSE")
	case true:
		fmt.Println("TRUE")
	}
}
