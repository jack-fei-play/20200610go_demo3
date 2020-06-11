package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	mongodb "github.com/jack-fei-play/20200610go_demo3/mongodb"
	mqttS "github.com/jack-fei-play/20200610go_demo3/mqtt_struct"
	"os"
)

const (
	ip       = "192.168.1.124"
	port     = 1883
	username = "admin"
	password = "public"
	//client := new (mqtt.client)
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	var mqttSend mqttS.MqttSendS
	err := json.Unmarshal(msg.Payload(), &mqttSend)
	if err != nil {
		fmt.Println("json convert fail,", err)
		return
	}

	//连接mongodb数据库
	mqttSendS := mongodb.FindMongodb(mqttSend)
	fmt.Println(mqttSendS)
	//mqtt发布消息
	MqttPublish(client, mqttSendS)

}

func main() {

	//mqtt连接emqx服务器，服务订阅者
	//获取mqtt连接信息对象ClientOptions
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", ip, port)).SetUsername(fmt.Sprintf("%s", username)).SetPassword(fmt.Sprintf("%s", password))

	//opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	//opts.SetPingTimeout(1 * time.Second)
	//获取mqtt连接对象Client
	c := mqtt.NewClient(opts)
	// 断开连接
	defer c.Disconnect(250)
	//client=c
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("mqtt listen 'ff/mqtt/message/get' topic!")
	for {
		// 订阅主题
		if token := c.Subscribe("ff/mqtt/message/get", 2, nil); token.Wait() && token.Error() != nil {
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
func MqttPublish(c mqtt.Client, mqttMongo mqttS.MqttMongodbS) {
	sendMsg, err := json.Marshal(mqttMongo)
	if err != nil {
		fmt.Println("json analysis fail,", err)
		return
	}
	//发布消息
	fmt.Println("mqtt publichser 'ff/mqtt/message/return' topic send:", string(sendMsg))
	token := c.Publish("ff/mqtt/message/return", 2, false, sendMsg)
	token.Wait()
}
