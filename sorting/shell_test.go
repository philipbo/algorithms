package main

import (
	"fmt"
	"github.com/philipbo/algorithms/utils"
	"testing"
)

func TestShellSort(t *testing.T) {
	arr := utils.RandArray(10)
	shellSort(arr)

	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println(arr)
			t.Error()
		}
	}
}

func benchmarkShellSort(n int, b *testing.B) {
	arr := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		shellSort(arr)
	}
}

func BenchmarkShellSort10(b *testing.B) {
	benchmarkShellSort(10, b)
}

func BenchmarkShellSort100(b *testing.B) {
	benchmarkShellSort(100, b)
}

func BenchmarkShellSort1000(b *testing.B) {
	benchmarkShellSort(1000, b)
}

func BenchmarkShellSort10000(b *testing.B) {
	benchmarkShellSort(10000, b)
}

func BenchmarkShellSort100000(b *testing.B) {
	benchmarkShellSort(100000, b)
}
