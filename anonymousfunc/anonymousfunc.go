package main

import "fmt"

func main() {
	for _, v := range []int{1, 2, 3, 4, 5} {
		go func(i int) {
			fmt.Println(i)
		}(v)
	}
}
