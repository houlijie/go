package main

import (
	"rabbitMq/direct"
)
//
func main()  {
	// direct.Producer()

	direct.Consumer()
}

// func hello(done chan bool) {
// 	fmt.Println( "Hello world goroutine")
// 	<-done
//
// }
// func main() {
// 	done := make(chan bool)
// 	done <- true
// 	go hello(done)
// 	fmt.Println("main function")
// 	time.Sleep(time.Second * 3)
// }