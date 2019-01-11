package main

/*
对Map使用“for range”语句迭代
如果你希望以某个顺序（比如，按key值排序）的方式得到元素，就需要这个技巧。每次的map迭代将会生成不同的结果。
Go的runtime有心尝试随机化迭代顺序，但并不总会成功，这样你可能得到一些相同的map迭代结果。
所以如果连续看到5个相同的迭代结果，不要惊讶。
*/

import "fmt"
func main() {
	m := map[string]int{"one":1,"two":2,"three":3,"four":4}
	for k,v := range m {
		fmt.Println(k,v)
	}
}

// 这个顺序是随机的