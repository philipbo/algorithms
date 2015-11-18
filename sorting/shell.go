package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
)

func shellSort(arr []int) {
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
}

func main() {
	arr := utils.RandArray(10)
	fmt.Println("[Shell] Unsorted array: ", arr)
	shellSort(arr)
	fmt.Println("[Shell] Sorted array: ", arr)
}
