package main

import (
	"fmt"
	"sync"
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

func main() {
	arr := []int{3, 1, 2, 7, 6, 8, 4, 9, 5}
	numWorkers := 4

	fmt.Println("Original array:", arr)

	parallelSelectionSort(arr, numWorkers)

	fmt.Println("Sorted array:", arr)
}
