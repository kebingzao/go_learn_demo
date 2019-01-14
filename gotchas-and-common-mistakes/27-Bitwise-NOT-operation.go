package main

import "fmt"

/*
按位NOT操作
许多语言使用 ~作为一元的NOT操作符（即按位补足），但Go为了这个重用了XOR操作符（^）, 也就是异或操作。
*/

//失败的例子：
/*func main() {
	fmt.Println(~2) //error
}*/


/*
print::
gotchas-and-common-mistakes\27-Bitwise-NOT-operation.go:12: syntax error: illegal character U+007E '~'
*/

// 正确的做法：
/*func main() {
	var d uint8 = 2
	fmt.Printf("%08b\n",^d)  // print 11111101
}*/

//Go依旧使用^作为XOR的操作符，这可能会让一些人迷惑。
//如果你愿意，你可以使用一个二元的XOR操作（如， 0x02 XOR 0xff）来表示一个一元的NOT操作（如，NOT 0x02）。
// 这可以解释为什么^被重用来表示一元的NOT操作。
//Go也有特殊的‘AND NOT’按位操作（&^），这也让NOT操作更加的让人迷惑。这看起来需要特殊的特性/hack来支持 A AND (NOT B)，而无需括号。

// 异或就是二进制对应的位如果不一样那么就是1，同样就是0. 同或如果都是1的话，才是1，其他都是0

func main() {
	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n",a)
	fmt.Printf("%08b [B]\n",b)
	fmt.Printf("%08b (NOT B)\n",^b)
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n",b,0xff,b ^ 0xff)
	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n",a,b,a ^ b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n",a,b,a & b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n",a,b,a &^ b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n",a,b,a & (^b))
}

// print::
/*10000010 [A]
00000010 [B]
11111101 (NOT B)
00000010 ^ 11111111 = 11111101 [B XOR 0xff]
10000010 ^ 00000010 = 10000000 [A XOR B]
10000010 & 00000010 = 00000010 [A AND B]
10000010 &^00000010 = 10000000 [A 'AND NOT' B]
10000010&(^00000010)= 10000000 [A AND (NOT B)]*/
