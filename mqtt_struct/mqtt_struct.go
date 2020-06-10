package mqtt_struct

type MqttSendS struct { //mqtt发送消息结构体
	DeviceId int32 `json:device_id`
}

type MqttReturnS struct { //mqtt返回消息结构体
}

type MqttMongodbS struct {
	ThisTime string             `json:"this_time"`
	DeviceId string             `json:"device_id"`
	Tag      string             `json:"tag"`
	Datas    []MqttMongodbDataS `json:"datas"`
}
type MqttMongodbDataS struct {
	TagId int     `json:"tag_id"`
	Value float32 `json:"value"`
}
