package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 3, 3, 6, 9, 9}, 4},
		{[]int{2, 3, 4, 5, 6, 9, 9}, 6},
		{[]int{2, 2, 2, 11}, 2},
		{[]int{1, 2, 3, 11, 11}, 4},
	}

	for _, tc := range testCases {
		actual := RemoveDuplicates(tc.input)
		if actual != tc.expected {
			t.Errorf("Input %v, expected %d, but got %d", tc.input, tc.expected, actual)
		}
	}
}
