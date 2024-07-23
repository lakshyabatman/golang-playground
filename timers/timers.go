package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	done := make(chan int)
	go func() {
		<-timer.C
		done <- 1
	}()
	// This way im making a go routine sync using a channel!
	<-done
	fmt.Println("Timer has ended:")

}
