package main

/*
被Defer的函数调用执行
被defer的调用会在包含的函数的末尾执行，而不是包含代码块的末尾。
对于Go新手而言，一个很常犯的错误就是无法区分被defer的代码执行规则和变量作用规则。
如果你有一个长时运行的函数，而函数内有一个for循环试图在每次迭代时都defer资源清理调用，那就会出现问题。
*/
/*
import (
	"fmt"
	"os"
	"path/filepath"
)
func main() {
	if len(os.Args) != 2 {
		os.Exit(-1)
	}
	start, err := os.Stat(os.Args[1])
	if err != nil || !start.IsDir(){
		os.Exit(-1)
	}
	var targets []string
	filepath.Walk(os.Args[1], func(fpath string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fi.Mode().IsRegular() {
			return nil
		}
		targets = append(targets,fpath)
		return nil
	})
	for _,target := range targets {
		f, err := os.Open(target)
		if err != nil {
			fmt.Println("bad target:",target,"error:",err) //prints error: too many open files
			break
		}
		defer f.Close() //will not be closed at the end of this code block， 这个只是代码块结尾，根本不会执行，只有函数末尾才会执行
		//do something with the file...
	}
}
*/

// 解决这个问题的一个方法是把代码块写成一个函数。

import (
	"fmt"
	"os"
	"path/filepath"
)
func main() {
	if len(os.Args) != 2 {
		os.Exit(-1)
	}
	start, err := os.Stat(os.Args[1])
	if err != nil || !start.IsDir(){
		os.Exit(-1)
	}
	var targets []string
	filepath.Walk(os.Args[1], func(fpath string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fi.Mode().IsRegular() {
			return nil
		}
		targets = append(targets,fpath)
		return nil
	})
	for _,target := range targets {
		func() {
			f, err := os.Open(target)
			if err != nil {
				fmt.Println("bad target:",target,"error:",err)
				return
			}
			defer f.Close() //ok
			//do something with the file...
		}()
	}
}

// 当然，最简单的方式就是去掉defer语句 :-)


