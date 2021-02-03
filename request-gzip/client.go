package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"flag"
)

func main() {
	var requestUseGzip = flag.String("req_use_gzip", "1", "has request body use gzip")
	var responseUseGzip = flag.String("resp_use_gzip", "1", "has response body use gzip")
	flag.Parse()
	var requestBodyStr = `{"name":"zach ke","des":"woooooo boy", "response_use_gzip": "%s"}`
	data := []byte(fmt.Sprintf(requestBodyStr, *responseUseGzip))
	var httpRequest *http.Request
	var err error
	if *requestUseGzip == "1"{
		fmt.Println("request with gzip")
		var zBuf bytes.Buffer
		zw := gzip.NewWriter(&zBuf)
		if _, err := zw.Write(data); err != nil {
			fmt.Println("gzip is faild,err:", err)
		}
		zw.Close()
		httpRequest, err = http.NewRequest("POST", "http://localhost:9905/request", &zBuf)
		if err != nil {
			fmt.Println("http request is failed, err: ", err)
		}
		httpRequest.Header.Set("Content-Encoding", "gzip")
	} else {
		fmt.Println("request without gzip")
		reader := bytes.NewReader(data)
		httpRequest, err = http.NewRequest("POST", "http://localhost:9905/request", reader)
		if err != nil {
			fmt.Println("http request is failed, err: ", err)
		}
	}
	// 通过参数判断是否返回值要用 gzip 压缩
	if *responseUseGzip == "1" {
		httpRequest.Header.Set("Accept-Encoding", "gzip")
	}else{
		// 这个也要指定，不然 response 那边获取 content-length 头部会取不到值
		httpRequest.Header.Set("Accept-Encoding", "deflate")
	}
	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println("httpResponse is failed, err: ", err)
	}
	defer httpResponse.Body.Close()

	fmt.Println("respone content length=", httpResponse.ContentLength)
	if httpResponse.StatusCode == 200 {
		var respBody string
		switch httpResponse.Header.Get("Content-Encoding") {
		case "gzip":
			fmt.Println("response with gzip")
			reader, err := gzip.NewReader(httpResponse.Body)
			if err != nil {
				fmt.Println("gzip get reader err ", err)
			}
			data, err = ioutil.ReadAll(reader)
			respBody = string(data)
		default:
			fmt.Println("response without gzip")
			bodyByte, _ := ioutil.ReadAll(httpResponse.Body)
			respBody = string(bodyByte)
		}
		fmt.Println("resp data=", respBody)
	}
}
