package mqtt_struct

type MqttSendS struct { //mqtt发送消息结构体
	DeviceId string `json:device_id`
}

type MqttResultS struct { //mqtt返回消息结构体
}

type MqttMongodbS struct { //post请求实体封装
	ThisTime string             `json:"this_time"`
	DeviceId string             `json:"device_id"`
	Tag      string             `json:"tag"`
	Datas    []MqttMongodbDataS `json:"datas"`
}
type MqttMongodbDataS struct {
	TagId int     `json:"tag_id"`
	Value float32 `json:"value"`
}
type MongoResultS struct { //mongo返回消息结构体
	Result  int    //返回响应状态  200
	Message string //返回信息描述  success
}
