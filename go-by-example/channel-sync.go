package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			fmt.Printf("%d Gorouting is working...\n", i)
			time.Sleep(300 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println("all done")
}
