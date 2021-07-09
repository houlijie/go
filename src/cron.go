package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func main()  {
	c := cron.New()
	_, err := c.AddFunc("* * * * *", func() {
		fmt.Println("hi, it's me!")
	})
	if err != nil {
		fmt.Println("add func failed, err: ", err)
	}

	c.Start()
	defer c.Stop()

	select {

	}
}
