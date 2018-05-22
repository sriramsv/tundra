package main

import (
	"github.com/sriramsv/tundra/mqtt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)


type ConnectHandler func()error

const SuccessStatus="SUCCESS"
const FailureStatus="FAILURE"
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
	Status string
	Error string
}
func NewHandler(mqttbroker,username,password,clientID string) *Handler{
	mclient:=mqtt.New()
	connection:=mclient.GetConnectionHandler(mqttbroker,clientID,username,password)
	err:=connection()
	if err!=nil{
		log.Fatalln("V:",err)
	}
	handler:=&Handler{
		mqttclient:mclient,
		ConnectHandler:connection,
	}
	return handler
}

func (h *Handler) PublishHandler(w http.ResponseWriter, r *http.Request){
	var p Payload;
	defer r.Body.Close()
	if r.Method!="POST"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(body,&p)
	if err!=nil{
		w.Write([]byte(err.Error()))
		return
	}
	resp:=*h.publishhandler(&p)
	log.Println(resp)
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)

}


func logHandler(p Payload){
	log.Println(fmt.Sprintf("Received message:(%s) on topic:(%s) [QOS:%d,retain:%v]",p.Message,p.Topic,p.Qos,p.Retain))
}

func (h *Handler) publishhandler( p *Payload)(*Response){
	payload:=*p
	logHandler(payload)
	resp:=new(Response)
	if payload.Topic==""{
		resp.Error="topic is empty"
		resp.Status=FailureStatus
		return resp
	}
	if payload.Message==""{
		resp.Error="message is empty"
		resp.Status=FailureStatus
		return resp
	}

	err:=h.mqttclient.Publish(p.Topic,p.Message,p.Retain,p.Qos)
	if err!=nil{
		resp.Error=err.Error()
		resp.Status=FailureStatus
	} else{
		resp.Error=""
		resp.Status=SuccessStatus
	}
	return resp

}