package main

import (
	"github.com/zouyx/agollo"
	"fmt"
	"iGong/json"
)

//func loadApolloConfig() {
//	readyConfig := &agollo.AppConfig{
//		AppId:         "dev-kbz-test",
//		NamespaceName: "application",
//		Ip:            "59.57.13.156:8889",
//	}
//
//	agollo.InitCustomConfig(func() (*agollo.AppConfig, error) {
//		return readyConfig, nil
//	})
//
//	config := agollo.GetAppConfig(nil)
//
//	fmt.Println(config.AppId)
//	fmt.Println(config.Cluster)
//	fmt.Println(config.NamespaceName)
//	fmt.Println(config.Ip)
//
//	apolloConfig := agollo.GetCurrentApolloConfig()
//	fmt.Println(apolloConfig.AppId)
//	fmt.Println(apolloConfig.Cluster)
//	fmt.Println(apolloConfig.NamespaceName)
//}

func main() {
	//loadApolloConfig()
	agollo.Start()
	// 监听状态改变通知
	event := agollo.ListenChangeEvent()
	go func(){
		for {
			select {
			case changeEvent := <-event:
				bytes, _ := json.Marshal(changeEvent)
				fmt.Println("event:", string(bytes))
				fmt.Println("namespace:", changeEvent.Namespace)
				// changeType: 0 是添加， 1 是修改， 2 为删除
				for k,v := range changeEvent.Changes {
					fmt.Println(fmt.Sprintf("val: %v, old: %v, new: %v, changeType: %v", k, v.OldValue, v.NewValue, v.ChangeType))
				}
			}
		}
	}()
	// 获取该应用的配置项
	age := agollo.GetIntValue("age", 12)
	fmt.Println("age:", age)
	name := agollo.GetStringValue("name", "default")
	fmt.Println("name:", name)
	select {}
}
