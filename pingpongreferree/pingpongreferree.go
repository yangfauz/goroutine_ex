package main

import (
	"fmt"
	"math/rand"
	"time"
)

//pingpong
//2 pemain pingpong
//permainan berhenti setelah 1 detik
//gunakan konkurensi dengan 1 channel

//tambahkan wasit
//wasit mengambil bola ketika salah satu pemain tidak dapat mengembalikan bola ke lawan
//pemain selesai, wasit tentukan pemenang

type Ball struct {
	hits       int
	lastPlayer string
}

func main() {
	table := make(chan *Ball)
	done := make(chan *Ball)
	go player("ping", table, done)
	go player("pong", table, done)

	referree(table, done)
}

func referree(table chan *Ball, done chan *Ball) {

	table <- new(Ball)

	for {
		select {
		case ball := <-done:
			fmt.Println("winner is", ball.lastPlayer)
			return
		}
	}
}

func player(name string, table chan *Ball, done chan *Ball) {
	for {

		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		select {
		case ball := <-table:
			v := r.Intn(1000)
			if v%11 == 0 {
				fmt.Println(name, "drop the ball")
				done <- ball
				return
			}

			ball.hits++
			ball.lastPlayer = name
			fmt.Println(name, "hits the ball", ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- ball

		case <-time.After(2 * time.Second):
			return
		}
	}
}
