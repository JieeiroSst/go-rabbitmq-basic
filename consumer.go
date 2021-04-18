package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main(){
	conn,err:=amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err!=nil {
		panic(err)
	}
	defer conn.Close()
	ch,err:=conn.Channel()
	if err!=nil {
		panic(err)
	}
	msgs,err:=ch.Consume(
		"Queue",
		"",
		true,
		false,
		false,
		false,
		nil,)
	if err!=nil {
		panic(err)
	}
	forver:=make(chan bool)
	go func() {
		for d:=range msgs {
			fmt.Printf("message %s \n",d.Body)
		}
	}()
	<-forver
}
