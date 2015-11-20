package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := utils.RandArray(10)
	mergeSort(arr)

	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println(arr)
			t.Error()
		}
	}
}

func benchmarkMergeSort(n int, b *testing.B) {
	arr := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		mergeSort(arr)
	}
}

func BenchmarkMergeSort10(b *testing.B) {
	benchmarkMergeSort(10, b)
}

func BenchmarkMergeSort100(b *testing.B) {
	benchmarkMergeSort(100, b)
}

func BenchmarkMergeSort1000(b *testing.B) {
	benchmarkMergeSort(1000, b)
}

func BenchmarkMergeSort10000(b *testing.B) {
	benchmarkMergeSort(10000, b)
}

func BenchmarkMergeSort100000(b *testing.B) {
	benchmarkMergeSort(100000, b)
}
