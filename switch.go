package main

import (
	"fmt"
	"runtime"
)



func main() {
	fmt.Println("GO runs on")
	switch os:=runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println(os)
	}
}