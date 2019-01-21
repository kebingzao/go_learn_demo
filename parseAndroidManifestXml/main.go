package main

import (
	"encoding/xml"
	"iGong/util/log"
	"github.com/shogo82148/androidbinary/apk"
	"fmt"
	"os"
	"github.com/shogo82148/androidbinary"
	"io/ioutil"
	"iGong/json"
)
// 解析apk包
func TestParseApk() {
	apkData, err := apk.OpenFile("./parseAndroidManifestXml/testdata/helloworld.apk")
	defer apkData.Close()
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println("apk NAME:", apkData.PackageName())
}
// 解析 manifest xml 文件
func TestParseManifestXml() {
	f, _ := os.Open("./parseAndroidManifestXml/testdata/AndroidManifest.xml")
	xmlFile, _ := androidbinary.NewXMLFile(f)
	reader := xmlFile.Reader()

	// read XML from reader
	var manifest apk.Manifest
	data, _ := ioutil.ReadAll(reader)
	fmt.Println("xml string:", string(data))
	xml.Unmarshal(data,manifest)
	str,_ := json.Marshal(manifest)
	// 将格式转换解析为 json格式
	fmt.Println("xml json data: ", string(str))
}
// 解析 arsc 文件
func TestParseResourceFile() {
	f, _ := os.Open("./parseAndroidManifestXml/testdata/resources.arsc")
	rsc, _ := androidbinary.NewTableFile(f)
	val, _ := rsc.GetResource(androidbinary.ResID(0x7f040000), nil)
	fmt.Println("resouce:", val)
}


func main(){
	TestParseApk()
	TestParseManifestXml()
	TestParseResourceFile()
}