package validation

import "testing"

// TestIsValidMobile tests the IsValidMobile function
func TestIsValidMobile(t *testing.T) {
	tests := []struct {
		mobile   int
		expected bool
	}{
		// Valid mobile numbers
		{1234567890, true},
		{9876543210, true},

		// Invalid mobile numbers
		{123456789, false},      // Less than 10 digits
		{12345678901, false},    // More than 10 digits
		{12345678, false},       // Less than 10 digits
		{12345678901234, false}, // More than 10 digits
		{0, false},              // Edge case: Single digit
		{-1234567890, false},    // Negative number
	}

	for _, test := range tests {
		result := IsValidMobile(test.mobile)
		if result != test.expected {
			t.Errorf("IsValidMobile(%d) = %v; expected %v", test.mobile, result, test.expected)
		}
	}
}
