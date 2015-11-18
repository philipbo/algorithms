package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := utils.RandArray(10)
	insertionSort(arr)

	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println(arr)
			t.Error()
		}
	}
}

func benchmarkInsertionSort(n int, b *testing.B) {
	arr := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		insertionSort(arr)
	}
}

func BenchmarkInsertionSort10(b *testing.B) {
	benchmarkInsertionSort(10, b)
}

func BenchmarkInsertionSort100(b *testing.B) {
	benchmarkInsertionSort(100, b)
}

func BenchmarkInsertionSort1000(b *testing.B) {
	benchmarkInsertionSort(1000, b)
}

func BenchmarkInsertionSort10000(b *testing.B) {
	benchmarkInsertionSort(10000, b)
}

func BenchmarkInsertionSort100000(b *testing.B) {
	benchmarkInsertionSort(100000, b)
}
