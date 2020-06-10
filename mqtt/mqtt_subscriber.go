package main

import (
	"../mqtt_struct"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

const (
	ip ="192.168.1.124"
	port =1883
	username="admin"
	password="public"
)
var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	var mqttSend mqtt_struct.MqttSendS
	err := json.Unmarshal(msg.Payload(), &mqttSend)
	if err != nil {
		fmt.Println("json convert fail,",err)
		return
	}

	//连接mongodb数据库

}

func main()  {

	//mqtt连接emqx服务器，服务订阅者
	//获取mqtt连接信息对象ClientOptions
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", ip, port)).SetUsername(fmt.Sprintf("%s", username)).SetPassword(fmt.Sprintf("%s", password))

	//opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	//opts.SetPingTimeout(1 * time.Second)
	//获取mqtt连接对象Client
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	for{
		// 订阅主题
		if token := c.Subscribe("ff/mqtt/message/1", 2, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}
	// 发布消息
	//token := c.Publish("testtopic/1", 0, false, "Hello World")
	//token.Wait()

	// 断开连接
	//c.Disconnect(250)
	//time.Sleep(30 * time.Second)
}


