package main

import (
	"math/rand"
	"sort"
	"testing"
)

func randlist(n int) []int {
	l := make([]int, n)
	for i := range l {
		l[i] = i
	}

	for i := range l {
		j := rand.Intn(i + 1)
		l[i], l[j] = l[j], l[i]
	}

	return l
}

func BenchmarkSortReflect(b *testing.B) {
	master := randlist(25)
	l := make([]int, 25)

	for i := 0; i < b.N; i++ {
		copy(l, master)
		sort.Sort(NewSortwrap(l, func(i, j int) bool {
			return l[i] < l[j]
		}))
	}
}

func BenchmarkSortRaw(b *testing.B) {
	master := randlist(25)
	l := make([]int, 25)

	for i := 0; i < b.N; i++ {
		copy(l, master)
		sort.Ints(l)
	}
}
