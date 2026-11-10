package problems

import (
	"fmt"
	"testing"
)

func TestFindMaxTableDriven(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 4, 3}, 4},
		{[]int{3, 3, 3}, 3},
		{[]int{-10, -2, -5}, -2},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.input)
		t.Run(testname, func(t *testing.T) {
			got := FindMax(tt.input)

			if got != tt.want {
				t.Errorf("FindMax(%v) = %d, want %d",
					tt.input, got, tt.want)
			}
		})
	}
}
