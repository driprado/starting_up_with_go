// https://golang.org/pkg/log/
// Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output. It also has a predefined 'standard' Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and Panic[f|ln]. That logger writes to standard error and prints the date and time of each logged message. Every log message is output on a separate line: if the message being printed does not end in a newline, the logger will add one. The Fatal functions call os.Exit(1) after writing the log message. The Panic functions call panic after writing the log message.

// // Example 1 - Basic Logging
// package main

// import "log"

// func main() {
// 	log.Println("Some log message") // You don't need to instantiate log and by default it includes date and time to stdout
// }

// // Example 2 - SetFlags
// // https://golang.org/pkg/log/#Logger.SetFlags This function lets you add more info to your log by adding optional flags in https://golang.org/pkg/log/#pkg-constants
// package main

// import "log"

// func main() {
// 	log.SetFlags(log.Lshortfile | log.LstdFlags | log.LUTC) // Include file that is generating the log and line number, print date and time in UTC
// 	log.Println("Some log message")
// }

// // Example 3 - log.Fatal
// // https://golang.org/pkg/log/#Fatal
// // Fatal is equivalent to Print() followed by a call to os.Exit(1).
// package main

// import "log"

// func main() {
// 	log.SetFlags(log.Lshortfile | log.LstdFlags | log.LUTC)
// 	log.Fatal("Some log message") // exit status 1
// }

// // Example 4 - Panic
// // Panic is equivalent to Print() followed by a call to panic().
// package main

// import "log"

// func main() {
// 	log.SetFlags(log.Lshortfile | log.LstdFlags | log.LUTC)
// 	log.Panic("Some log message") // Logs message and then generates panic in the go routine
// }

// Example 5 - Writing Log into location
// https://golang.org/pkg/log/#Logger
package main

import (
	"log"
	"os"
)

var (
	// We will create specializes loggers of type log, which will be instantiated later on
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLoger    *log.Logger
)

func main() {
	file, err := os.OpenFile("mylog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // Using os package to create a file with permission 0666
	// Log Fatal error in case expression above fails
	if err != nil {
		log.Fatal(err)
	}
	// Initializes loggers
	infoLogger = log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)
	warningLogger = log.New(file, "WARNING: ", log.LstdFlags|log.Lshortfile)
	errorLoger = log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)

	// Loggers write to file
	infoLogger.Println("Informational log about your app")
	warningLogger.Println("I'm warning you!")
	errorLoger.Println("Error! this is a error!")

}
