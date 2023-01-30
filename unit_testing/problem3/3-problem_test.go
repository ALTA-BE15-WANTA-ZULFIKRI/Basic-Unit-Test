package main

import (
	"fmt"
	"testing"
)

func TestMunculSekali(t *testing.T) {
	var tests = []struct {
		input string
		want  []int
	}{
		{"1234123", []int{4}},
		{"76523752", []int{6, 3}},
		{"12345", []int{1, 2, 3, 4, 5}},
		{"1122334455", []int{}},
		{"0872504", []int{8, 7, 2, 5, 4}},
	}

	for _, test := range tests {
		if got := munculsekali(test.input); fmt.Sprint(got) != fmt.Sprint(test.want) {
			t.Errorf("munculsekali(%q) = %v", test.input, got)
		}
	}
}