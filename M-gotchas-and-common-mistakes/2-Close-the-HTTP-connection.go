package main

/*
关闭HTTP的连接
一些HTTP服务器保持会保持一段时间的网络连接（根据HTTP 1.1的说明和服务器端的“keep-alive”配置）。
默认情况下，标准http库只在目标HTTP服务器要求关闭时才会关闭网络连接。
这意味着你的应用在某些条件下消耗完sockets/file的描述符。

你可以通过设置请求变量中的Close域的值为true，来让http库在请求完成时关闭连接。

另一个选项是添加一个Connection的请求头，并设置为close。
目标HTTP服务器应该也会响应一个Connection: close的头。当http库看到这个响应头时，它也将会关闭连接。
*/

//例子：
/*
import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func main() {
	req, err := http.NewRequest("GET","http://golang.org",nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Close = true
	//or do this:
	//req.Header.Add("Connection", "close")
	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(string(body)))
}
*/


//你也可以取消http的全局连接复用。你将需要为此创建一个自定义的http传输配置。

import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func main() {
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://golang.org")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(string(body)))
}

/*
如果你向同一个HTTP服务器发送大量的请求，那么把保持网络连接的打开是没问题的。
然而，如果你的应用在短时间内向大量不同的HTTP服务器发送一两个请求，那么在引用收到响应后立刻关闭网络连接是一个好主意。
增加打开文件的限制数可能也是个好主意。当然，正确的选择源自于应用。*/
