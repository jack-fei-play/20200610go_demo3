package main

import (
	"encoding/json"
	"fmt"
	mongodb "github.com/jack-fei-play/20200610go_demo3/mongodb"
	mqS "github.com/jack-fei-play/20200610go_demo3/mqtt_struct"
	"io/ioutil"
	"net/http"
)

//http
func main() {

	http.HandleFunc("/message", msgPost)
	fmt.Println("Starting 8080 server ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server 8080 start fail!", err)
	}
}

func msgPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		println("json:", string(body))
		var mqttMongodbS mqS.MqttMongodbS
		err = json.Unmarshal(body, &mqttMongodbS)
		if err != nil {
			fmt.Println(err)
			return
		}
		//向mongodb写入数据
		mongodb.CoonMongodb(mqttMongodbS)
		w.Write([]byte("Starting 8080 server"))
	}

}
