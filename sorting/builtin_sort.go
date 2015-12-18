package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
	"sort"
)

func builtinSort(arr []int) {
	sort.Ints(arr)
}

func main() {
	arr := utils.RandArray(10)
	fmt.Println("[Builtin Sort] Unsorted array: ", arr)
	builtinSort(arr)
	fmt.Println("[Builtin Sort] Sorted array: ", arr)

}
