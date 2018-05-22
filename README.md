# tundra
HTTP-MQTT bridge in pure golang :)


Tundra is a HTTP-MQTT bridge written in pure golang.

## Installation

```
$ go get github.com/sriramsv/tundra
$ go install github.com/sriramsv/tundra
```

## Usage

```

tundra --help

Usage: tundra [OPTIONS]

HTTP-MQTT bridge

Options:
      --port             port to listen on! (env $TUNDRA_PORT) (default "3000")
      --mqtt-user        username of mqtt broker (env $TUNDRA_MQTT_USER)
      --mqtt-port        mqtt broker port (env $TUNDRA_MQTT_PORT)
      --mqtt-pwd         mqtt broker password (env $TUNDRA_MQTT_PWD)
      --mqtt-host        MQTT broker host (env $TUNDRA_MQTT_HOST)
      --mqtt-client-id   port to listen on! (env $TUNDRA_MQTT_CLIENT_ID) (default "tundra")
      
```

TO publish a message to the broker, send a POST request to the server as follows:
```bash

curl -X POST \
  http://localhost:3000/publish \
  -H 'content-type: application/json' \
  -d '{
	"topic":"tundra",
	"message":"its alright!",
	"qos": 2,
	"retain": false
}'
```


Note: The connection from the server to the broker only supports tcp connections currently. So your data will go unencrypted over the wire.


## TODO
* Add TLS support
* Add websocket support for subscriptions to topics

