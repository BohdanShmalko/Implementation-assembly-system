package main

import (
	"testing"
)

func TestNumbers_MulSqrt(t *testing.T) {
	testCases := []struct {
		name            string
		numbers         *Numbers
		isExpectedError bool
	}{
		{
			name:            "two positive numbers",
			numbers:         NewNumbers(10, 20),
			isExpectedError: false,
		},
		{
			name:            "two negative numbers",
			numbers:         NewNumbers(-10, -20),
			isExpectedError: false,
		},
		{
			name:            "positive and negative numbers",
			numbers:         NewNumbers(10, -20),
			isExpectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.numbers.MulSqrt()
			if err != nil && !tc.isExpectedError {
				t.Error(err)
			}
		})
	}
}
