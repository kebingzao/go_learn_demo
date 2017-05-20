package main

import "os"
// panic 错误处理
// http://www.yiibai.com/go/golang-panic.html
func main() {

	// We'll use panic throughout this site to check for
	// unexpected errors. This is the only program on the
	// site designed to panic.
	panic("a problem")

	// A common use of panic is to abort if a function
	// returns an error value that we don't know how to
	// (or want to) handle. Here's an example of
	// `panic`king if we get an unexpected error when creating a new file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}