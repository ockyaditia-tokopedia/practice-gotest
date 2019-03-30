package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		TestName       string
		A              int
		B              int
		ExpectedResult int
	}{
		{"Neg", 1, 2, 3},
		{"Pos", 1, 2, 2},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			if add(test.A, test.B) != test.ExpectedResult {
				t.Errorf("Wrong output")
			}
		})
	}

	// for _, test := range tests {
	// 	if add(test.A, test.B) != test.ExpectedResult {
	// 		t.Errorf("Wrong output")
	// 	}
	// }

	a, b := 1, 2
	result := 2

	if add(a, b) != result {
		t.Errorf("Wrong output")
	}
}
