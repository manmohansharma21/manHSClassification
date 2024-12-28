package runner

import (
	"os/exec"
)

// RunPythonScript runs a Python script and returns any error
func RunPythonScript(scriptPath string) error {
	cmd := exec.Command("python3", scriptPath)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
