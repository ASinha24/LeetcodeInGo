package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		target  int
		indices []int
	}{
		{
			name:    "two nnumber sum should be equal to 9",
			nums:    []int{2, 7, 11, 15},
			target:  9,
			indices: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.indices) {
				t.Errorf("RotateArrayFromRight() expected %v got %v", tt.indices, got)
			}
		})
	}
}
