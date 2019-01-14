package main

/*
关闭HTTP的响应
当你使用标准http库发起请求时，你得到一个http的响应变量。如果你不读取响应主体，你依旧需要关闭它。
注意对于空的响应你也一定要这么做。对于新的Go开发者而言，这个很容易就会忘掉。

一些新的Go开发者确实尝试关闭响应主体，但他们在错误的地方做。
*/

//失败的例子：
/*import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	defer resp.Body.Close()//not ok
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}*/


//这段代码对于成功的请求没问题，但如果http的请求失败，resp变量可能会是nil，这将导致一个runtime panic。
//最常见的关闭响应主体的方法是在http响应的错误检查后调用defer。

//正确的是:

/*
import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()//ok, most of the time :-)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
*/

//大多数情况下，当你的http响应失败时，resp变量将为nil，而err变量将是non-nil。
// 然而，当你得到一个重定向的错误时，两个变量都将是non-nil。这意味着你最后依然会内存泄露。

//通过在http响应错误处理中添加一个关闭non-nil响应主体的的调用来修复这个问题。
// 另一个方法是使用一个defer调用来关闭所有失败和成功的请求的响应主体。

import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
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
	fmt.Println(string(body))
}


/*
resp.Body.Close()的原始实现也会读取并丢弃剩余的响应主体数据。
这确保了http的链接在keepalive http连接行为开启的情况下，可以被另一个请求复用。最新的http客户端的行为是不同的。
现在读取并丢弃剩余的响应数据是你的职责。如果你不这么做，http的连接可能会关闭，而无法被重用。这个小技巧应该会写在Go 1.5的文档中。

如果http连接的重用对你的应用很重要，你可能需要在响应处理逻辑的后面添加像下面的代码：
_, err = io.Copy(ioutil.Discard, resp.Body)

如果你不立即读取整个响应将是必要的，这可能在你处理json API响应时会发生：
json.NewDecoder(resp.Body).Decode(&data)
*/

