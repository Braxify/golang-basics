package main

import (
	"fmt"
	"sync"
	"time"
)

func selectionSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func merge(arr1, arr2 []int) []int {
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			merged[k] = arr1[i]
			i++
		} else {
			merged[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len(arr1) {
		merged[k] = arr1[i]
		i++
		k++
	}

	for j < len(arr2) {
		merged[k] = arr2[j]
		j++
		k++
	}

	return merged
}

func parallelSelectionSort(arr []int, numWorkers int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	chunkSize := (length + numWorkers - 1) / numWorkers
	chunks := make([][]int, numWorkers)
	wg := sync.WaitGroup{}
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > length {
			end = length
		}

		go func(index, start, end int) {
			defer wg.Done()
			chunk := arr[start:end]
			selectionSort(chunk)
			chunks[index] = chunk
		}(i, start, end)
	}

	wg.Wait()

	for i := 1; i < numWorkers; i++ {
		chunks[0] = merge(chunks[0], chunks[i])
	}

	copy(arr, chunks[0])
}

func measureTime(arr []int, numWorkers int) time.Duration {
	start := time.Now()
	parallelSelectionSort(arr, numWorkers)
	end := time.Now()
	return end.Sub(start)
}

func computeSpeedup(sequentialTime, parallelTime time.Duration) float64 {
	return float64(sequentialTime) / parallelTime.Seconds()
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

func main() {
	arr := []int{3, 1, 2, 7, 6, 8, 4, 9, 5}
	dimensions := []int{1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000}
	numWorkersList := []int{2, 4, 6, 8, 10, 20}

	for _, dimension := range dimensions {
		arr = generateRandomArray(dimension)

		sequentialTime := measureTime(make([]int, len(arr)), 1)
		fmt.Printf("Dimension: %d\n", dimension)

		for _, numWorkers := range numWorkersList {
			parallelTime := measureTime(arr, numWorkers)
			speedup := computeSpeedup(sequentialTime, parallelTime)

			fmt.Printf("NumWorkers: %d, Speedup: %f\n", numWorkers, speedup)
		}

		fmt.Println()
	}
}
