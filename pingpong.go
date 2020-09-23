package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func () {c <- "saque"}()

	go pong(c)
	go ping(c)
	go pong(c)
	go ping(c)
	go pong(c)
	go ping(c)
	time.Sleep(1 * time.Minute)
}

func ping(c chan string) {
	fmt.Println(<- c)
	c <- "ping"
}

func pong(c chan string) {
	fmt.Println(<- c)
	c <- "pong"
}