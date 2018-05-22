package main

import (

	"github.com/jawher/mow.cli"
	"net/http"
	"log"
	"os"
	"fmt"
)


func getAddr(host,port string)string{
	return fmt.Sprintf("%s:%s",host,port)
}
func main() {
	hostname,_:=os.Hostname()
	log.Println(hostname)
	app := cli.App("tundra","HTTP-MQTT bridge")
	cmd := app.Cmd
	port := cmd.String(cli.StringOpt{
		Name:   "port",
		Desc:   "port to listen on!",
		Value:  "3000",
		EnvVar: "TUNDRA_PORT",
	})
	mqtt_user := cmd.String(cli.StringOpt{
		Name:   "mqtt-user",
		Desc:   "username of mqtt broker",
		Value:  "",
		EnvVar: "TUNDRA_MQTT_USER",
	})
	mqtt_port := cmd.String(cli.StringOpt{
		Name:   "mqtt-port",
		Desc:   "mqtt broker port",
		Value:  "",
		EnvVar: "TUNDRA_MQTT_PORT",
	})
	mqtt_pwd := cmd.String(cli.StringOpt{
		Name:   "mqtt-pwd",
		Desc:   "mqtt broker password",
		Value:  "",
		EnvVar: "TUNDRA_MQTT_PWD",
	})
	mqtt_host := cmd.String(cli.StringOpt{
		Name:   "mqtt-host",
		Desc:   "MQTT broker host",
		Value:  "",
		EnvVar: "TUNDRA_MQTT_HOST",
	})
	mqtt_clientId := cmd.String(cli.StringOpt{
		Name:   "mqtt-client-id",
		Desc:   "mqtt client id",
		Value:  hostname,
		EnvVar: "TUNDRA_MQTT_USER",
	})

	cmd.Action = func() {

			mux := http.NewServeMux()
			if *mqtt_host==""{
				log.Fatal("No Broker host URL specified")
			}
			if *mqtt_port==""{
				log.Fatal("No Broker Port specified")
			}
			address:=getAddr(*mqtt_host,*mqtt_port)
			h:=NewHandler(address,*mqtt_user,*mqtt_pwd,*mqtt_clientId)
			ph := http.HandlerFunc(h.PublishHandler)
			mux.Handle("/publish", ph)
			log.Println("Listening on port",*port)
			addr:=fmt.Sprintf("0.0.0.0:%s",*port)
			http.ListenAndServe(addr, mux)
		}
	app.Run(os.Args)
}