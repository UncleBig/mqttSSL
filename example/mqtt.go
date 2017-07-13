package main

import (
	"charles/mqttSSL"
	"fmt"
	"time"
)

func main() {
	mqttSSL.Init()

	mqttSSL.MqttSSLCli.Subscribe("/go-mqtt/sample", 0, nil)

	i := 0
	for _ = range time.Tick(time.Duration(1) * time.Second) {
		if i == 5 {
			break
		}
		text := fmt.Sprintf("this is msg #%d!", i)
		mqttSSL.MqttSSLCli.Publish("/go-mqtt/sample", 0, false, text)
		i++
	}

	mqttSSL.MqttSSLCli.Disconnect(250)

}
