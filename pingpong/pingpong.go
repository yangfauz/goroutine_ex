package main

import (
	"fmt"
	"time"
)

//pingpong
//2 pemain pingpong
//permainan berhenti setelah 1 detik
//gunakan konkurensi dengan 1 channel

type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) //game on; toss the ball
	time.Sleep(1 * time.Second)
	<-table // game over; grab the ball
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
