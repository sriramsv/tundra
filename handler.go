package main

import (
	"github.com/sriramsv/tundra/mqtt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"go.corp.yahoo.com/clusterville/log"
	"fmt"
)


type ConnectHandler func()error

type  Handler struct{
	mqttclient *mqtt.MQTTClient
	ConnectHandler ConnectHandler
}

type Payload struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
	Qos     int    `json:"qos"`
	Retain  bool `json:"retain"`
}

type Response struct {
	Payload Payload
	Status string
	Error error
}
func NewHandler(mqttbroker,username,password,clientID string) *Handler{
	mclient:=mqtt.New()
	connection:=mclient.GetConnectionHandler(mqttbroker,clientID,username,password)
	err:=connection()
	if err!=nil{
		log.Fatalln(err)
	}
	handler:=&Handler{
		mqttclient:mclient,
		ConnectHandler:connection,
	}
	return handler
}

func (h *Handler) PublishHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var p Payload
	body, rerr := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if rerr!=nil{
		log.Println(rerr)
	}
	err := json.Unmarshal(body,&p)
	if err!=nil{
		log.Println(err)
	}
	logHandler(p)
	err=h.mqttclient.Publish(p.Topic,p.Message,p.Retain,p.Qos)
	resp:=new(Response)
	resp.Payload=p
	if err!=nil{
		resp.Error=err
		resp.Status="FAILURE"
	} else{
		resp.Error=nil
		resp.Status="SUCCESS"
	}
	json.NewEncoder(w).Encode(resp)
}


func logHandler(p Payload){
	log.Println(fmt.Sprintf("Received message:(%s) on topic:(%s) [QOS:%d,retain:%v]",p.Message,p.Topic,p.Qos,p.Retain))
}