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
	fmt.Println("Successfully connected to our rabbitmq instance")

	ch,err:=conn.Channel()
	if err!=nil {
		panic(err)
	}
	defer ch.Close()

	q,err:=ch.QueueDeclare(
		"Queue",
		false,
		false,
		false,
		false,
		nil,
		)
	if err!=nil {
		panic(err)
	}
	fmt.Println(q)

	err=ch.Publish(
		"",
		"Queue",
		false,
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[] byte("hello world"),
		})
	if err!=nil {
		panic(err)
	}

}