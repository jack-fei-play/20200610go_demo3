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
	url = "mongodb://192.168.200.129:27017"
)

func CoonMongodb(s mqS.MqttMongodbS) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(url)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("mymongo").Collection("ffmqttdemo")
	insertResult, err := collection.InsertOne(context.TODO(), s)
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
