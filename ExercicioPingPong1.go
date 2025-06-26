package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Ping"
			time.Sleep(time.Second * 1)
		}

	}()

	go func() {
		for {
			c2 <- "Pong"
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}

}
