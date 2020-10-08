// // TODO
// https://blog.golang.org/using-go-modules <= Readup/Writeup

// 1 - Create hello.go

// 2 - Create hello_test.go

// 3 - Run go test in current directory // should pass

// 4 - Initialize module
// $ go mod init hello // should create go.mod file

// 5 - Test your module
// $ go test   // should pass and print the name of the module:

// 6 - Adding a dependency https://blog.golang.org/using-go-modules

package hello

// Hello returns greeting
func Hello() string {
	return "Hello, world."
}
