package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
)

func insertionSort(arr []int) {
	j := 0
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		for j = i; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

func main() {
	arr := utils.RandArray(10)
	fmt.Println("[Insertion] Unsorted array: ", arr)
	insertionSort(arr)
	fmt.Println("[Insertion] Sorted array: ", arr)
}
