package direct

import (
	"log"

	"rabbitMq/util"
)

func Consumer()  {
	conn := util.ConnOpen()
	defer util.ConnClose(conn)

	ch := util.ChanOpen(conn)
	defer util.ChanClose(ch)

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		util.FailOnError(err, "声明队列失败")
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		util.FailOnError(err, "获取消息失败")
	}


	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %+v", d)
			// ch.Reject(d.DeliveryTag, true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	// <-forever



}
