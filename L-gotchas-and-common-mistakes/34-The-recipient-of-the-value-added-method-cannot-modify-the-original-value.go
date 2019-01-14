package main

/*
传值方法的接收者无法修改原有的值
方法的接收者就像常规的函数参数。
如果声明为值，那么你的函数/方法得到的是接收者参数的拷贝。
这意味着对接收者所做的修改将不会影响原有的值，除非接收者是一个map或者slice变量，而你更新了集合中的元素，或者你更新的域的接收者是指针。
*/


import "fmt"
type data struct {
	num int
	key *string
	items map[string]bool
}
func (this *data) pmethod() {
	this.num = 7
}
func (this data) vmethod() {
	this.num = 8
	*this.key = "v.key"
	this.items["vmethod"] = true
}
func main() {
	key := "key.1"
	d := data{1,&key,make(map[string]bool)}
	fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	//prints num=1 key=key.1 items=map[]
	d.pmethod()
	fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	//prints num=7 key=key.1 items=map[]
	d.vmethod()
	fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	//prints num=7 key=v.key items=map[vmethod:true]
	// items 会改变是因为他是map，本来就是引用传递，之所以 num 是 7，是因为调用 vmethod 赋值 num 的 data 对象不是指针，而是 值传递，所以不会影响原来的值
}