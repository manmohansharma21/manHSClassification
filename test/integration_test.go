package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestGoAndPythonIntegration(t *testing.T) {
	// Step 1: Generate sample data using the Python script
	cmd := exec.Command("python", "scripts/create_sample_data.py")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to generate sample data: %v", err)
	}

	// Step 2: Process the data using the Python script
	cmd = exec.Command("python", "scripts/process_data.py")
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to process data: %v", err)
	}

	// Step 3: Run the Go application to categorize the descriptions
	cmd = exec.Command("go", "run", "cmd/main.go")
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run Go application: %v", err)
	}

	// Step 4: Verify that the output file is generated
	if _, err := os.Stat("data/processed_descriptions.xlsx"); os.IsNotExist(err) {
		t.Fatalf("Output file processed_descriptions.xlsx does not exist")
	}
}
