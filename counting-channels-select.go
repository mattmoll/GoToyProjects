package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Every second 900k messages are sent on Whatsapp")
	fmt.Println("Every 5 seconds Bill Gates earns U$D 1.250")

	fmt.Println("Lets visualize it with go.")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			c1 <- "900k messages sent on Whatsapp"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			c2 <- "Bill gates earns U$D 1.250"
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