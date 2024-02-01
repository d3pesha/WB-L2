package main

import "testing"

var tests = []struct {
	input  string
	output string
	err    bool
}{
	{"a4bc2d5e", "aaaabccddddde", false},
	{"abcd", "abcd", false},
	{"45", "", true},
	{"", "", false},
}

func TestUnpack(t *testing.T) {
	for _, test := range tests {
		result, err := Unpack(test.input)
		if (result != test.output) || ((err != nil) != test.err) {
			t.Errorf("Unpack %s = %s, expected %s, error: %t, expected %t",
				test.input, result, test.output, err != nil, test.err)
		}
	}
}
