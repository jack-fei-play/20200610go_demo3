package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	mqttS "github.com/jack-fei-play/20200610go_demo3/mqtt_struct"
	"time"
)

func main() {

	//mqtt连接emqx服务器，服务订阅者
	//获取mqtt连接信息对象ClientOptions
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", "192.168.1.124", 1883)).SetUsername(fmt.Sprintf("%s", "admin")).SetPassword(fmt.Sprintf("%s", "public"))

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetPingTimeout(1 * time.Second)
	//获取mqtt连接对象Client
	c := mqtt.NewClient(opts)
	// 断开连接
	defer c.Disconnect(250)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	var mqttSendS mqttS.MqttSendS
	mqttSendS.DeviceId = "897979799"
	sendMsg, err := json.Marshal(mqttSendS)
	if err != nil {
		fmt.Println("json analysis fail,", err)
		return
	}
	fmt.Println("mqtt publichser 'ff/mqtt/message/get' topic send:", string(sendMsg))
	//发布消息
	token := c.Publish("ff/mqtt/message/get", 0, false, sendMsg)
	token.Wait()

}
