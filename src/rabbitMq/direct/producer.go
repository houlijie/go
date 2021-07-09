package direct

import (
	"github.com/streadway/amqp"
	"rabbitMq/util"
)

func main()  {
	Producer()
}

func Producer()  {
	conn := util.ConnOpen()
	defer util.ConnClose(conn)

	ch := util.ChanOpen(conn)
	defer util.ChanClose(ch)

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	util.FailOnError(err, "创建队列失败")

	msg := "hello world!"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(msg),
	})
	util.FailOnError(err, "消息创建失败")
}

