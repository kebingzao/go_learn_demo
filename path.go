package main

import (
	"fmt"
	"path"
)

func main() {
	//  Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠。如果路径是空字符串，会返回"."；如果路径由1到多个斜杠后跟0到多个非斜杠字符组成，会返回"/"；其他任何情况下都不会返回以斜杠结尾的路径。
	fmt.Println(" :", path.Dir(""))
	fmt.Println(". :", path.Dir("."))
	fmt.Println("a :", path.Dir("a"))
	fmt.Println("/ :", path.Dir("/"))
	fmt.Println("/a :", path.Dir("/a"))
	fmt.Println("/a/b/ :", path.Dir("/a/b/"))
	fmt.Println("/a/b :", path.Dir("/a/b"))
	fmt.Println("a/b :", path.Dir("a/b"))
	fmt.Println("a/b/ :", path.Dir("a/b/"))
	fmt.Println("/////// :", path.Dir("///////"))
	fmt.Println("///////a :", path.Dir("///////a"))
	fmt.Println("====================================================")
	//  Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠。结果是经过简化的，所有的空字符串元素会被忽略。
	fmt.Println(", :", path.Join("", ""))
	fmt.Println("/, :", path.Join("/", ""))
	fmt.Println(", / :", path.Join("", "/"))
	fmt.Println("/, / :", path.Join("/", "/"))
	fmt.Println("/a, :", path.Join("/a", ""))
	fmt.Println(", /a :", path.Join("", "/a"))
	fmt.Println("/a, / :", path.Join("/a", "/"))
	fmt.Println("/, /a :", path.Join("/", "/a"))
	fmt.Println("a, b :", path.Join("a", "b"))
	fmt.Println("/a, /b :", path.Join("/a", "/b"))
	fmt.Println(", ////// :", path.Join("", "//////"))
	fmt.Println("//////,  :", path.Join("//////", ""))
	fmt.Println(", /////a :", path.Join("", "/////a"))
	fmt.Println("/////a, :", path.Join("/////a", ""))
	fmt.Println("/a, /////b :", path.Join("/a", "/////b"))
	fmt.Println("a, /////b :", path.Join("a", "/////b"))
	fmt.Println("/////a, /b :", path.Join("/////a", "/b"))
}