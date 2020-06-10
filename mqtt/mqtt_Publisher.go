package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)


func main()  {

	//mqtt连接emqx服务器，服务订阅者
	//获取mqtt连接信息对象ClientOptions
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", "192.168.1.124", 1883)).SetUsername(fmt.Sprintf("%s", "admin")).SetPassword(fmt.Sprintf("%s", "public"))

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetPingTimeout(1 * time.Second)
	//获取mqtt连接对象Client
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	//发布消息
	token := c.Publish("ff/mqtt/message/1", 2, false, "Hello World")
	token.Wait()

	// 断开连接
	c.Disconnect(250)
}


