package main

/*
未导出的结构体不会被编码
以小写字母开头的结构体将不会被（json、xml、gob等）编码，因此当你编码这些未导出的结构体时，你将会得到零值。
*/

import (
	"fmt"
	"encoding/json"
)
type MyData struct {
	One int
	two string
}
func main() {
	in := MyData{1,"two"}
	fmt.Printf("%#v\n",in) //prints main.MyData{One:1, two:"two"}
	encoded,_ := json.Marshal(in)
	fmt.Println(string(encoded)) //prints {"One":1}
	var out MyData
	json.Unmarshal(encoded,&out)
	fmt.Printf("%#v\n",out) //prints main.MyData{One:1, two:""}
}