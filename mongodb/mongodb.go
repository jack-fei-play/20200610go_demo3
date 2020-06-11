package mongodb

import (
	"context"
	"fmt"
	mqS "github.com/jack-fei-play/20200610go_demo3/mqtt_struct"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	url = "mongodb://192.168.1.124:27017"
)

func CoonMongodb(s mqS.MqttMongodbS) bool {
	// 设置客户端连接配置
	var credential options.Credential
	credential.Password = "123456"
	credential.Username = "cyf"
	clientOptions := options.Client().SetAuth(credential).ApplyURI(url)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("mymongo").Collection("ffmqttdemo")
	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}
func FindMongodb(s mqS.MqttSendS) mqS.MqttMongodbS {
	// 设置客户端连接配置
	var credential options.Credential
	credential.Password = "123456"
	credential.Username = "cyf"
	clientOptions := options.Client().SetAuth(credential).ApplyURI(url)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return mqS.MqttMongodbS{}
	}
	defer client.Disconnect(context.TODO())

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return mqS.MqttMongodbS{}
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("mymongo").Collection("ffmqttdemo")
	cursor, err := collection.Find(context.TODO(), s)
	if err != nil {
		fmt.Println(err)
		return mqS.MqttMongodbS{}
	}
	var result *mqS.MqttMongodbS
	for cursor.Next(context.TODO()) {
		//创建需要反序列化成什么样子的结构体对象
		//反序列化
		mqMonS := &mqS.MqttMongodbS{}
		err = cursor.Decode(mqMonS)
		if err != nil {
			fmt.Println(err)
			return mqS.MqttMongodbS{}
		}
		//打印
		result = mqMonS
		//fmt.Println(*result)
	}
	return *result
}
