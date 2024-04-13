package intersectionoftwoarray

import (
	"reflect"
	"testing"
)

func TestIntersectionOfArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "two nnumber sum should be equal to 9",
			nums1: []int{2, 7, 2, 15},
			nums2: []int{2, 2},
			want:  []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntersectionOfArrays(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateArrayFromRight() expected %v got %v", tt.want, got)
			}
		})
	}
}
