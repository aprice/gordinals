package ordinal

import (
	"testing"
	"fmt"
)

func TestFor(t *testing.T) {
	tests := []struct{
		in int
		expected string
	}{
		{0, "th"},
		{1, "st"},
		{2, "nd"},
		{3, "rd"},
		{4, "th"},
		{11, "th"},
		{12, "th"},
		{13, "th"},
		{21, "st"},
		{101, "st"},
		{111, "th"},
		{1011, "th"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("n=%d",test.in), func(tt *testing.T) {
			actual := For(test.in)
			if actual != test.expected {
				tt.Errorf("For(%v): expected %s, got %s", test.in, test.expected, actual)
			}
		})
	}
}

func BenchmarkFor(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			For(111)
		}
	})
}
