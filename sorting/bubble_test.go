package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := utils.RandArray(10)
	bubbleSort(arr)

	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println(arr)
			t.Error()
		}
	}
}

func benchmarkBubbleSort(n int, b *testing.B) {
	arr := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		bubbleSort(arr)
	}
}

func BenchmarkBubbleSort10(b *testing.B) {
	benchmarkBubbleSort(10, b)
}

func BenchmarkBubbleSort100(b *testing.B) {
	benchmarkBubbleSort(100, b)
}

func BenchmarkBubbleSort1000(b *testing.B) {
	benchmarkBubbleSort(1000, b)
}

func BenchmarkBubbleSort10000(b *testing.B) {
	benchmarkBubbleSort(10000, b)
}

func BenchmarkBubbleSort100000(b *testing.B) {
	benchmarkBubbleSort(100000, b)
}
