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

		j := i

		go func() {
			defer wg.Done()

			fmt.Printf("%d Gorouting is working...\n", j)
			time.Sleep(300 * time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Println("all done")
}
