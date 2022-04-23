package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	blocker := make(chan struct{})
	<-blocker
}

func producer(ch chan int) {
	counter := 0

	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		<-ticker.C
		ch <- counter
		counter++
	}
}

func consumer(ch chan int) {
	for {
		counter := <-ch
		fmt.Println(counter)
	}
}
