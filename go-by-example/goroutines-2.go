package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go test(ch)

	for i := range ch {
		fmt.Println(i)
	}
}

func test(ch chan int) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- i
	}

	close(ch)
}
