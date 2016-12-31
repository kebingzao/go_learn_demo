package main

import "fmt"
//import "math"
//func swap (x, y string)(string, string){
//	return y, x
//}
//func split(sum int)(x, y int)  {
//	x = sum * 4/9
//	y = sum -x
//	return
//}

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt (x int) int {
	return x * 10 + 1
}

func needFloat (x float64) float64 {
	return x * 0.1
}

func main() {

	//fmt.Printf("Hello, world")
	//fmt.Println(math.Pi)
	//a,b := swap("kk","bb")
	//fmt.Println(a,b)
	//fmt.Println(split(17))


	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))
}