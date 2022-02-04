package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		<-time.After(time.Duration(rand.Intn(2)) * time.Second)
		c <- 10
	}()

	select {
	case val := <-c:
		fmt.Println(val)
	case <-time.After(time.Duration(rand.Intn(2)) * time.Second):
		fmt.Println("timeout")
	}
}
