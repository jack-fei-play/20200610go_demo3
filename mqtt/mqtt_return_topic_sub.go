package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

const (
	ip1       = "192.168.1.124"
	port1     = 1883
	username1 = "admin"
	password1 = "public"
	//client := new (mqtt.client)
)

var g mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

}

func main() {
	//mqtt连接emqx服务器，服务订阅者
	//获取mqtt连接信息对象ClientOptions
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", ip1, port1)).SetUsername(fmt.Sprintf("%s", username1)).SetPassword(fmt.Sprintf("%s", password1))

	//opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(g)
	//opts.SetPingTimeout(1 * time.Second)
	//获取mqtt连接对象Client
	c := mqtt.NewClient(opts)
	// 断开连接
	defer c.Disconnect(250)
	//client=c
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("mqtt listen 'ff/mqtt/message/return' topic!")
	for {
		// 订阅主题
		if token := c.Subscribe("ff/mqtt/message/return", 2, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}
}
