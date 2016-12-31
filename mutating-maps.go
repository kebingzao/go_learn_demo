package main

import (
	"fmt"
)




func main() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("the value:", m["answer"])

	m["answer"] = 48
	fmt.Println("the value:", m["answer"])

	//delete(m, "answer")
	//fmt.Println("the value:", m["answer"])

	v, ok := m["answer"]
	fmt.Println("the value:", v, "present?", ok)
}