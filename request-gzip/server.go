package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func handler(resp http.ResponseWriter, req *http.Request) {
	var bodyDataStr string
	// 是否需要解 gzip 压缩
	fmt.Println("resquest content length=", req.ContentLength)
	if req.Header.Get("Content-Encoding") == "gzip" {
		fmt.Println("request with gzip")
		body, err := gzip.NewReader(req.Body)
		if err != nil {
			fmt.Println("unzip is failed, err:", err)
		}
		defer body.Close()
		data, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println("read all is failed.err:", err)
		}
		bodyDataStr = string(data)
	} else {
		fmt.Println("request without gzip")
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			fmt.Println("read resp is failed, err: ", err)
		}
		bodyDataStr = string(data)
	}
	fmt.Println("request json string=", bodyDataStr)

	respJson := []byte(`{code: "1",msg: "success"}`)

	// 通过头部的 Accept-Encoding 判断返回值是否要用 gzip 压缩
	if req.Header.Get("Accept-Encoding") == "gzip" {
		fmt.Println("response with gzip")
		// 添加 gzip 头部
		resp.Header().Set("Content-Encoding", "gzip")
		var zBuf bytes.Buffer
		zw := gzip.NewWriter(&zBuf)
		if _, err := zw.Write(respJson); err != nil {
			zw.Close()
			fmt.Println("gzip is faild,err:", err)
		}
		zw.Close()
		//fmt.Println("gzip content:", string(zBuf.Bytes()))
		resp.Write(zBuf.Bytes())
	}else{
		fmt.Println("response without gzip")
		// 正常不压缩 返回
		resp.Write(respJson)
	}
}

func main() {
	fmt.Println("http://localhost:9905/request")
	http.HandleFunc("/request", handler)
	http.ListenAndServe(":9905", nil)
}
