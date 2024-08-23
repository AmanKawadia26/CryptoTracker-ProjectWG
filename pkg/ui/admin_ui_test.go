package ui

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/fatih/color"
)

// Helper function to simulate input and capture output
func simulateAdminPanelInput(inputs []string) string {
	// Create a buffer to capture output
	var buf bytes.Buffer

	// Create a pipe for input
	inputReader, inputWriter, _ := os.Pipe()

	// Backup original stdin and stdout
	oldStdin := os.Stdin
	oldStdout := os.Stdout

	// Replace stdin with our pipe
	os.Stdin = inputReader

	// Replace stdout with our buffer
	os.Stdout = &buf

	// Disable color output for testing
	color.NoColor = true

	// Write inputs to the input pipe
	go func() {
		for _, input := range inputs {
			fmt.Fprintln(inputWriter, input)
		}
		inputWriter.Close()
	}()

	// Run the admin panel function
	ShowAdminPanel()

	// Restore original stdin and stdout
	os.Stdin = oldStdin
	os.Stdout = oldStdout

	// Close the input reader
	inputReader.Close()

	// Return captured output
	return buf.String()
}

func Test_ShowAdminPanel(t *testing.T) {
	// Test case for choosing "Logout" (option 4)
	output := simulateAdminPanelInput([]string{"4"})

	// Print the captured output for debugging
	fmt.Printf("Captured output:\n%s\n", output)

	expectedItems := []string{
		"Admin Panel",
		"1. Manage Users",
		"2. View User Profiles",
		"3. Manage User Requests",
		"4. Logout",
		"Enter your choice:",
		"Logging out...",
	}

	// Check if output contains all expected items
	for _, item := range expectedItems {
		if !strings.Contains(output, item) {
			t.Errorf("Expected output to contain '%s', but it was not found.", item)
		}
	}
}
