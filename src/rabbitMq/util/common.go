package util

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnOpen() *amqp.Connection  {
	conn, err:= amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "创建连接失败")
	return conn
}

func ConnClose(conn *amqp.Connection)   {
	err := conn.Close()
	FailOnError(err, "连接关闭失败")
}

func ChanOpen(conn *amqp.Connection) *amqp.Channel  {
	ch, err := conn.Channel()
	FailOnError(err, "创建管道失败")
	return ch
}

func ChanClose(ch *amqp.Channel)  {
	err := ch.Close()
	FailOnError(err, "管道关闭失败")
}

func FailOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("$s: %s", msg, err)
	}
}
