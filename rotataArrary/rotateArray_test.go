package rotataarrary

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	tests := []struct {
		name          string
		array         []int
		expectedArray []int
	}{
		{
			name:          "reverse the array",
			array:         []int{1, 2, 3, 4, 5},
			expectedArray: []int{5, 4, 3, 2, 1},
		},
		{
			name:          "reverse the array",
			array:         []int{5, 4, 3, 2, 1},
			expectedArray: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseArray(tt.array, 0, len(tt.array)-1)
			if !reflect.DeepEqual(tt.array, tt.expectedArray) {
				t.Errorf("ReverseArray() expected %v go %v", tt.expectedArray, tt.array)
			}
		})
	}
}

func TestRotateArrayFromRight(t *testing.T) {
	tests := []struct {
		name          string
		array         []int
		kTimes        int
		expectedArray []int
	}{
		{
			name:          "reverse the array",
			array:         []int{1, 2, 3, 4, 5},
			kTimes:        3,
			expectedArray: []int{3, 4, 5, 1, 2},
		},
		{
			name:          "reverse the array",
			array:         []int{5, 7, 9, 4, 6, 7, 3, 8, 9, 4, 8},
			kTimes:        3,
			expectedArray: []int{9, 4, 8, 5, 7, 9, 4, 6, 7, 3, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RotateArrayFromRight(tt.array, tt.kTimes)
			if !reflect.DeepEqual(got, tt.expectedArray) {
				t.Errorf("RotateArrayFromRight() expected %v got %v", tt.expectedArray, tt.array)
			}
		})
	}
}

func TestRotateArrayFromLeft(t *testing.T) {
	tests := []struct {
		name          string
		array         []int
		kTimes        int
		expectedArray []int
	}{
		{
			name:          "reverse the array",
			array:         []int{1, 2, 3, 4, 5},
			kTimes:        3,
			expectedArray: []int{4, 5, 1, 2, 3},
		},
		{
			name:          "reverse the array",
			array:         []int{5, 7, 9, 4, 6, 7, 3, 8, 9, 4, 8},
			kTimes:        3,
			expectedArray: []int{4, 6, 7, 3, 8, 9, 4, 8, 5, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RotateArrayFromLeft(tt.array, tt.kTimes)
			if !reflect.DeepEqual(got, tt.expectedArray) {
				t.Errorf("RotateArrayFromRight() expected %v got %v", tt.expectedArray, tt.array)
			}
		})
	}
}
