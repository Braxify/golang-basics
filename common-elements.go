package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{5, 2}

	commonElements := make(map[int]bool)

	var mutex sync.Mutex
	var wg sync.WaitGroup

	wg.Add(len(arr1))

	for _, val1 := range arr1 {
		go func(val1 int) {
			for _, val2 := range arr2 {
				if val1 == val2 {
					mutex.Lock()
					commonElements[val1] = true
					mutex.Unlock()
					break
				}
			}
			wg.Done()
		}(val1)
	}

	// Очікуємо завершення роботи всіх горутин
	wg.Wait()

	fmt.Println("Кількість спільних елементів:", len(commonElements))

	if len(commonElements) > 0 {
		fmt.Print("Спільні елементи: ")
		for val := range commonElements {
			fmt.Print(val, " ")
		}
	}

	fmt.Println()

	endTime := time.Now()
	duration := endTime.Sub(startTime).Seconds()
	// seconds := float64(duration.Nanoseconds()) / 1e9

	fmt.Printf("Час виконання програми: %.6f с\n", duration)
}
