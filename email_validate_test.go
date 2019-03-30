package main

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		TestName       string
		Email          string
		ExpectedResult bool
	}{
		{"Test Case 1", "ja.ne.doe@tokopedia.com", true},
		{"Test Case 2", "doe@doe.com", true},
		{"Test Case 3", "12A3213__123@numbers.com", true},
		{"Test Case 4", "tokopedia.com", false},
		{"Test Case 5", "john.smith@tokopedia", true},
		{"Test Case 6", "jack.sparrow", false},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			if validateEmail(test.Email) != test.ExpectedResult {
				t.Fatalf("Wrong Test Case")
			}
		})
	}
}
