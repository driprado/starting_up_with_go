package main

import "fmt"

//// Reading
// https://www.youtube.com/watch?v=i3o4G4bmqPc 			@done
// https://www.youtube.com/watch?v=93f9_bJQdHk			@done
// https://golang.org/ref/spec#Method_declarations	@done
// https://golang.org/ref/spec#Method_sets
// https://golang.org/doc/faq#methods_on_values_or_pointers

//// Notes:
// A method is a function with a receiver.
// A method declaration associates the method with the receiver (ex: struct)

// There are two types of methods in Go:
// - Pointer receivers
// 		.	capable of changing values in the receiver's base type (ex: structs)
// - Value receivers
// 		. perform operations with receiver values, but doesn't permanently change the receiver's base type
// Example 1: Pointer receiver (modify values of the base type, in this case, the 'trip' struct)
type trip struct {
	distance        float64
	time            float64
	speed           float64
	fuelConsumption float64
}

// (t *trip)	 	method receiver, bounds the method to the trip type
// calcTime()	 	method name
// float64	 		return type
func (t *trip) calcTime() float64 {
	t.time = t.distance / t.speed
	return t.time
}

func main() {
	// toDarwing is passing values to the trip struct
	toDarwing := &trip{
		distance:        3942,
		time:            0,
		speed:           120,
		fuelConsumption: 0,
	}
	fmt.Println(toDarwing.time) // time value is current 0
	toDarwing.calcTime()        // execute the calcTime() method with the values of 'toDarwing'
	fmt.Println(toDarwing.time) // after method is applied, time is now 32.85, this changed a value in the struct, rather than just making calculations externally
}
