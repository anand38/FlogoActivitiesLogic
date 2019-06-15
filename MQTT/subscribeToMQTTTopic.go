package mqttTrigger

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
	"syscall"
)
var activityLog = logger.GetLogger("trigger-flogo-MQTTT")
var knt int
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
	fmt.Sprintf("this is result msg #%d!", knt)

	//text := fmt.Sprintf("this is result msg #%d!", knt)
	knt++
	//token := client.Publish("nn/result", 0, false, text)
	//token.Wait()
}
func SubscribeToMQTTTopic(url string,clientID string, mqttTopic string){
	knt = 0
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := MQTT.NewClientOptions().AddBroker(url)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(f)
	topic := mqttTopic

	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		activityLog.Info("Connected to server\n")
		//fmt.Printf("Connected to server\n")
	}
	<-c
}


func main()  {
	subscribeToMQTTTopic("tcp://localhost:1883","sub","test")
}
 

