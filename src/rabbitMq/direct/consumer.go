package direct

import (
	"fmt"
	"log"
)

func Consumer()  {
	// conn := util.ConnOpen()
	// defer util.ConnClose(conn)
	//
	// ch := util.ChanOpen(conn)
	// defer util.ChanClose(ch)
	//
	// q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	// if err != nil {
	// 	util.FailOnError(err, "声明队列失败")
	// }
	//
	// _, err = ch.Consume(q.Name, "", true, false, false, false, nil)
	// if err != nil {
	// 	util.FailOnError(err, "获取消息失败")
	// }

	forever := make(chan bool)

	forever <- true

	select {
	case aa := <-forever:
		fmt.Println(aa)
	default:
		fmt.Println("ssssss")
	}
	
	// go func() {
	// 	for d:= range msgs {
	// 		forever <- true
	// 		log.Printf("Received a message: %s", d.Body)
	// 	}
	// }()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// <-forever



}
