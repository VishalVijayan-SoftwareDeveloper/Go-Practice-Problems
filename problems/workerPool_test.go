package problems

import (
	"slices"
	"testing"
)

func TestSquareReadings(t *testing.T) {
	var tests = []struct {
		input       []int
		workerCount int
		want        []int
	}{
		{[]int{2, 3, 4}, 4, []int{4, 9, 16}},
	}

	for _, tt := range tests {

		t.Run("", func(t *testing.T) {
			got := SquareReadings(tt.input, tt.workerCount)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
