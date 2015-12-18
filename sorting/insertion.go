package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp
	}
}

func main() {
	arr := utils.RandArray(10)
	fmt.Println("[Insertion] Unsorted array: ", arr)
	insertionSort(arr)
	fmt.Println("[Insertion] Sorted array: ", arr)
}
