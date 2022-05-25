package main

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

func Test_FibRecursive_Can_Compute_The_Nth_Fibonacci_Number(t *testing.T) {
	t.Run("works for small numbers", func(t *testing.T) {
		knownFibs := map[int]int{}
		assert.Equal(t, 0, FibRecursive(0, knownFibs))
		assert.Equal(t, 3, FibRecursive(4, knownFibs))
		assert.Equal(t, 610, FibRecursive(15, knownFibs))
	})

	t.Run("works for bigger numbers", func(t *testing.T) {
		knownFibs := map[int]int{}
		assert.Equal(t, 7540113804746346429, FibRecursive(92, knownFibs))
	})

	t.Run("panics when value becomes too big for int64", func(t *testing.T) {
		knownFibs := map[int]int{}
		assert.Panics(t, func() {
			FibRecursive(93, knownFibs)
		})
	})
}

func Test_FibIterative_Can_Compute_The_Nth_Fibonacci_Number(t *testing.T) {
	t.Run("works for small numbers", func(t *testing.T) {
		assert.Equal(t, 0, FibIterative(0))
		assert.Equal(t, 3, FibIterative(4))
		assert.Equal(t, 610, FibIterative(15))
	})

	t.Run("works for bigger numbers", func(t *testing.T) {
		assert.Equal(t, 7540113804746346429, FibIterative(92))
	})

	t.Run("panics when value becomes too big for int64", func(t *testing.T) {
		assert.Panics(t, func() {
			FibIterative(93)
		})
	})
}

func BenchmarkFibRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		knownFibs := map[int]int{}
		for n := 0; n <= 92; n++ {
			FibRecursive(n, knownFibs)
		}
	}
}

func BenchmarkFibIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := 0; n <= 92; n++ {
			FibIterative(n)
		}
	}
}
