package binarysearch

import (
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		list []int
		item int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 6, false},
		{[]int{}, 3, false},
	}

	for _, tt := range tests {
		got := main(tt.list, tt.item)
		if got != tt.want {
			t.Errorf("main(%v, %d) = %v; want %v", tt.list, tt.item, got, tt.want)
		}
	}
}

func Benchmark(b *testing.B) {
	list := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = i
	}
	item := 500

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		main(list, item)
	}
}
