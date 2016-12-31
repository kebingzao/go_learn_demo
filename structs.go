package main

import "fmt"
type Vertex struct {
	X int
	Y int
}

var (
	v1 =Vertex{1,2}
	v2 = Vertex{X:1}
	v3 = Vertex{}
	p4 = &Vertex{1,2}
)

func main() {
	//v := Vertex{4,7}
	//fmt.Println(v)
	//
	//v.X = 9
	//fmt.Println(v.X)
	//
	//p := &v
	//p.X = 1e9
	//fmt.Println(v)


	fmt.Println(v1, v2, v3, p4)

}