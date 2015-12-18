package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
	"testing"
)

func TestBuiltinSort(t *testing.T) {
	arr := utils.RandArray(10)
	builtinSort(arr)

	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println(arr)
			t.Error()
		}
	}
}

func benchmarkBuiltinSort(n int, b *testing.B) {
	arr := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		builtinSort(arr)
	}
}

func BenchmarkBuiltinSort10(b *testing.B) {
	benchmarkBuiltinSort(10, b)
}

func BenchmarkBuiltinSort100(b *testing.B) {
	benchmarkBuiltinSort(100, b)
}

func BenchmarkBuiltinSort1000(b *testing.B) {
	benchmarkBuiltinSort(1000, b)
}

func BenchmarkBuiltinSort10000(b *testing.B) {
	benchmarkBuiltinSort(10000, b)
}

func BenchmarkBuiltinSort100000(b *testing.B) {
	benchmarkBuiltinSort(100000, b)
}
