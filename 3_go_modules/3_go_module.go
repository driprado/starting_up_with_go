// https://blog.golang.org/using-go-modules

// 1 - Create hello.go

// 2 - Create hello_test.go

// 3 - Run go test in current directory // should pass

// 4 - Initialize module
// $ go mod init hello // should create go.mod file

// 5 - Test your module
// $ go test   // should pass and print the name of the module:

// package hello

// // Hello returns greeting
// func Hello() string {
// 	return "Hello, world."
// }

// 6 - Adding a dependency https://blog.golang.org/using-go-modules

package hello

import "rsc.io/quote"

// Hello returns a quote
func Hello() string {
	return quote.Hello()
}

// 6.1 Test code again
// go.mod should've changed and added new "rsc.io/quote" dependency

// 6.2 Run "go list -m all" on your code's directory to check your dependencies
// $ go list -m all
// hello
// golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
// rsc.io/quote v1.5.2
// rsc.io/sampler v1.3.0

// 7 - Upgrading dependencies
// $ go get golang.org/x/text
// go: golang.org/x/text upgrade => v0.3.3
