package main

import (
	"github.com/kebingzao/agollo"
	"fmt"
	"iGong/json"
	"iGong/util/log"
)

func main() {
	log.Info("====start")
	// 启动 start
	start()
	// 获取该应用的配置项
	getStringValue()
	// 获取配置中所有的key
	getAllKeys()
	go watchUpdate()
	go subNamespace()
	log.Info("====end")
	select {}
}
func start(){
	// 这个会读取默认的 app.properties
	//if err := agollo.Start(); err != nil {
	//	log.Errorf("Start with default app.properties should return err, got :%v", err)
	//}
	// 当然我们可以自己拼凑 Conf 的格式：
	conf := &agollo.Conf{
		AppID:  "dev-kbz-test",
		Cluster: "default",
		NameSpaceNames : []string{"application","TEST1.kbz-namespace-test"},
		IP  : "59.57.13.156:8060",
	}
	if err := agollo.StartWithConf(conf); err != nil {
		log.Errorf("Start with default app.properties should return err, got :%v", err)
	}

}
func watchUpdate()  {
	// 监听状态改变通知
	event := agollo.WatchUpdate()
	for {
		select {
		case changeEvent := <-event:
			bytes, _ := json.Marshal(changeEvent)
			log.Info("event:", string(bytes))
			log.Info("namespace:", changeEvent.Namespace)
			// changeType: 0 是添加， 1 是修改， 2 为删除
			for k,v := range changeEvent.Changes {
				log.Info(fmt.Sprintf("val: %v, old: %v, new: %v, changeType: %v", k, v.OldValue, v.NewValue, v.ChangeType))
			}
		}
	}
}

func getStringValue()  {
	age := agollo.GetStringValue("age", "12")
	log.Info("age:", age)
	gname := agollo.GetStringValueWithNameSpace("TEST1.kbz-namespace-test", "name", "name1")
	log.Info("gname:", gname)
	// 获取 content 字段
	content := agollo.GetNameSpaceContent("TEST1.kbz-namespace-test", "default content")
	log.Info("content:", content)
}
func getAllKeys(){
	keys := agollo.GetAllKeys("application")
	log.Info("key length:", len(keys))
	log.Info("all key with application:", keys)
	keys2 := agollo.GetAllKeys("TEST1.kbz-namespace-test")
	log.Info("key2 length:", len(keys2))
	log.Info("all key2 with application:", keys2)
}
func subNamespace(){
	agollo.SubscribeToNamespaces("application", "TEST1.kbz-namespace-test")
}