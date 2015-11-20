package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
)

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
}

func main() {
	arr := utils.RandArray(10)
	fmt.Println("[Merge] Unsorted array: ", arr)
	mergeSort(arr)
	fmt.Println("[Merge] Sorted array: ", arr)
}
