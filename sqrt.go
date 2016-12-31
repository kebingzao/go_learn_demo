package main

import "fmt"

func Sqrt (x float64) float64 {
	z:= float64(1)
	for i:=1; i<=10; i++ {
		z = z -(z * z -x)/(2*x)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}