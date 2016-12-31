package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

//var m map[string]Vertex

var m = map[string]Vertex {
	"kbz gst": Vertex{ 40.111, 34.234},
	// 最后一个一定要有逗号，不然会报错
	"hello": Vertex{50.333,7.999},
}
var m2 = map[string]Vertex {
	"gst":  {6.111, 8.234},
	// 最后一个一定要有逗号，不然会报错
	"ho": {590.333,7.00999},
}

func main() {
	//m = make(map[string]Vertex)
	//m["Bell Labs"] = Vertex{
	//	40.68433, -74.39967,
	//}
	//
	//fmt.Println(m["Bell Labs"])
	fmt.Println(m)
	fmt.Println(m2)
}