package validation

import "testing"

// TestIsValidPassword tests the IsValidPassword function
func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		// Valid passwords
		{"P@ssw0rd", true},  // Minimum length, includes upper, lower, number, and special char
		{"A1b@cdef", true},  // Minimum length, includes upper, lower, number, and special char
		{"1234Abcd!", true}, // Minimum length, includes upper, lower, number, and special char

		// Invalid passwords
		{"short", false},          // Less than 8 characters
		{"NoSpecialChar1", false}, // Missing special character
		{"NOLOWER123!", false},    // Missing lowercase letter
		{"noupper123!", false},    // Missing uppercase letter
		{"NoNumber!", false},      // Missing number
		{"12345678", false},       // Missing upper, lower, and special char
		{"ABCD1234", false},       // Missing lowercase and special char
		{"abcd1234", false},       // Missing uppercase and special char
		{"ABCDabcd", false},       // Missing number and special char
		{"ABCD1234@", false},      // Missing lowercase letter
		{"1234!@#$", false},       // Missing uppercase and lowercase letters
	}

	for _, test := range tests {
		result := IsValidPassword(test.password)
		if result != test.expected {
			t.Errorf("IsValidPassword(%q) = %v; expected %v", test.password, result, test.expected)
		}
	}
}
