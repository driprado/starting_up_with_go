// // Reading:

// https://golang.org/ref/spec#Interface_types
// https://www.alexedwards.net/blog/interfaces-explained

// What are interfaces?
// interface is a type in Go
// An interface type specifies a method set called its interface.

// A simple File interface:
// interface {
// 	Read([]byte) (int, error)
// 	Write([]byte) (int, error)
// 	Close() error
// }

// An interface T may use a interface type name E in place of a method specification.
// This is called embedding interface E in T.
// The method set of T is the union of the method sets of T’s explicitly declared methods and of T’s embedded interfaces:
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// 	Close() error
// }
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// 	Close() error
// }
// // ReadWriter's methods are Read, Write, and Close.
// type ReadWriter interface {
// 	Reader  // includes methods of Reader in ReadWriter's method set
// 	Writer  // includes methods of Writer in ReadWriter's method set
// }