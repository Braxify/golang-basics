package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")
	go say("hi")
}

func say(w string) {
	// time.Sleep(5 * time.Second)
	fmt.Println(w)
}
