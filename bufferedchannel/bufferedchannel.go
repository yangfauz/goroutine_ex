package main

import "fmt"

func main() {
	c := make(chan bool, 5)
	c <- true

	fmt.Println("this line be printed")
}
