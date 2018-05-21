package mqtt

import (
	mclient "github.com/yosssi/gmq/mqtt/client"
	"log"
)


type MQTTClient struct{
	client *mclient.Client

}

func New() *MQTTClient{
	mcli:=new(MQTTClient)
	cli := mclient.New(&mclient.Options{
		ErrorHandler: func(err error) {
			log.Println(err.Error())
		},
	})
	mcli.client=cli
	return mcli
}



func (m *MQTTClient) GetConnectionHandler(address,clientID,username,password string)func()error{
	f:=func()error{
		return m.connect(address,clientID,username,password)
		}
		return f
}
func (m *MQTTClient) connect (address,clientID,username,password string)error{
	err := m.client.Connect(&mclient.ConnectOptions{
		Network:  "tcp",
		Address:  address,
		ClientID: []byte(clientID),
		UserName: []byte(username),
		Password: []byte(password),
		CleanSession: false,
		KeepAlive: 60,
	})
	return err
}

func (m *MQTTClient) Publish(topic string,message string,retain bool,qos int) error {

	err := m.client.Publish(&mclient.PublishOptions{
		QoS:       byte(qos),
		Retain:    retain,
		TopicName: []byte(topic),
		Message:   []byte(message),
	})
	return err
}
