# 20200610go_demo3
基于MQTT和API案例
项目介绍：通过API接口将数据保存到Mongodb中，并使用MQTT来实现数据的查询交互过程
##### 1.实现简单交互功能

​	1.1 对外提供post请求API接口，将数据保存到mongodb的（集合名称：ff-mqtt-demo）

```json
1.url: localhost:8080/message   method:post
参数：{  
        "this_time": "1581417203000",
        "device_id": "897979799",
        "tag": "dtu0000001",
        "datas": [ 
        {"tag_id":1101, "value":1.3},
        {"tag_id":1102, "value":25},
        {"tag_id":1103, "value":80}]
	}
返回结果：{"result":200,"message":"success"}
```

​	1.2 订阅ff/mqtt/message/get，并根据内容去mongodb中查询对应device_id的信息，将内容以主题ff/mqtt/message/return进行返回

```json
1.ff/mqtt/message/get主题内容
{"device_id":"897979799"}

2.ff/mqtt/message/return主题内容
{  
        "this_time": "1581417203000",
        "device_id": "897979799",
        "tag": "dtu0000001",
        "datas": [ 
        {"tag_id":1101, "value":1.3},
        {"tag_id":1102, "value":25},
        {"tag_id":1103, "value":80}]
}
```



