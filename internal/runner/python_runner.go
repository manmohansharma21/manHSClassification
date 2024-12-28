package runner

import (
	"fmt"
	"os/exec"
)

// RunPythonScript runs a Python script and returns any error
func RunPythonScript(scriptPath string) error {
	cmd := exec.Command("python3", scriptPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running script %s: %v\nOutput: %s", scriptPath, err, string(output))
	}
	fmt.Printf("Script %s executed successfully.\n", scriptPath)
	return nil
}
