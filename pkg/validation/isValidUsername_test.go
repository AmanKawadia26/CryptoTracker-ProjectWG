package validation

import "testing"

// TestIsValidUsername tests the IsValidUsername function
func TestIsValidUsername(t *testing.T) {
	tests := []struct {
		username string
		expected bool
	}{
		// Valid usernames
		{"user123", true},
		{"user_name", true},
		{"UserName", true},
		{"user", true},
		{"username_", true},
		{"_username", true},
		{"u123", true},

		// Invalid usernames
		{"user!123", false},     // Contains special character
		{"user name", false},    // Contains space
		{"user@name", false},    // Contains special character
		{"user-name", false},    // Contains hyphen
		{"user.name", false},    // Contains period
		{"", false},             // Empty username
		{"user name123", false}, // Contains space
	}

	for _, test := range tests {
		result := IsValidUsername(test.username)
		if result != test.expected {
			t.Errorf("IsValidUsername(%q) = %v; expected %v", test.username, result, test.expected)
		}
	}
}
