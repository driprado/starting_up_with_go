// Consolidate this:
// https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31
// https://deepsource.io/blog/go-modules/
// https://blog.golang.org/using-go-modules




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

// $ go  list -m all
// hello
// golang.org/x/text v0.3.3									// => text is updated and has new indirect dependency below
// golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
// rsc.io/quote v1.5.2
// rsc.io/sampler v1.3.0

// Make sure your code still works
// $ go test
// PASS
// ok      hello   0.006s

// updaintg sampler
// $ go get rsc.io/sampler
// go: rsc.io/sampler upgrade => v1.99.99

// $ go list -m all
// hello
// golang.org/x/text v0.3.3
// golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
// rsc.io/quote v1.5.2
// rsc.io/sampler v1.99.99 //==> Updated

// $ go test
// --- FAIL: TestHello (0.00s)
//     3_go_module_test.go:8: Hello() = "99 bottles of beer on the wall, 99 bottles of beer, ...", want "Hello, world."
// FAIL
// exit status 1
// FAIL    hello   0.005s //==> Fail, let's try anoter version of sampler

// $ go get rsc.io/sampler@v1.3.1

// $ go list -m all
// hello
// golang.org/x/text v0.3.3
// golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
// rsc.io/quote v1.5.2
// rsc.io/sampler v1.3.1 //==> feched version

// $ go test
// PASS						//>> We have a pass this time
// ok      hello   0.006s
