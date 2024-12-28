package runner

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

// Mocking exec.Command to test without actually running a Python script
func mockExecCommand(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	return cmd
}

// Test the RunPythonScript function
func TestRunPythonScript(t *testing.T) {
	// Save the original exec.Command function
	originalExecCommand := execCommand
	defer func() { execCommand = originalExecCommand }() // Restore the original function after the test

	// Mock exec.Command for testing purposes
	execCommand = mockExecCommand

	tests := []struct {
		name       string
		scriptPath string
		wantErr    bool
	}{
		{
			name:       "Valid Python script",
			scriptPath: "valid_script.py",
			wantErr:    false,
		},
		{
			name:       "Invalid Python script",
			scriptPath: "invalid_script.py",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RunPythonScript(tt.scriptPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunPythonScript() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Custom function for testing purposes, replacing the original exec.Command
var execCommand = exec.Command

// Helper function to check if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
