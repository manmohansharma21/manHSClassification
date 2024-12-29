// internal/hsclassification/hsclassification_test.go
package hsclassification

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRunner is a mock for the runner package's RunPythonScript function
type MockRunner struct {
	mock.Mock
}

func (m *MockRunner) RunPythonScript(scriptPath string) error {
	args := m.Called(scriptPath)
	return args.Error(0)
}

// TestRunHSCodeClassification tests the RunHSCodeClassification function
func TestRunHSCodeClassification(t *testing.T) {
	mockRunner := new(MockRunner)

	// Mocking the behavior of the RunPythonScript function
	mockRunner.On("RunPythonScript", "scripts/create_sample_data.py").Return(nil)
	mockRunner.On("RunPythonScript", "scripts/convert_to_csv.py").Return(nil)
	mockRunner.On("RunPythonScript", "scripts/process_data.py").Return(nil)

	// Replace the real runner with the mock
	// Since we're directly calling the functions from hsclassification package,
	// we can call the RunHSCodeClassification directly and inject the mockRunner.
	// If your functions depend on runner.RunPythonScript directly, you need to refactor them
	// to accept the runner interface as a dependency.

	query := "I want a laptop case"
	// We call the function that runs the classification
	RunHSCodeClassification(query)

	// Assert that all expectations were met
	mockRunner.AssertExpectations(t)
}

// TestGenerateTestData tests the GenerateTestData function
func TestGenerateTestData(t *testing.T) {
	mockRunner := new(MockRunner)

	// Mock the behavior of the RunPythonScript function
	mockRunner.On("RunPythonScript", "scripts/create_sample_data.py").Return(nil)

	// Call the function and assert it does not return an error
	err := GenerateTestData()
	assert.NoError(t, err)

	// Assert that the RunPythonScript function was called with the expected argument
	mockRunner.AssertExpectations(t)
}

// TestConvertToCSV tests the ConvertToCSV function
func TestConvertToCSV(t *testing.T) {
	mockRunner := new(MockRunner)

	// Mock the behavior of the RunPythonScript function
	mockRunner.On("RunPythonScript", "scripts/convert_to_csv.py").Return(nil)

	// Call the function and assert it does not return an error
	err := ConvertToCSV()
	assert.NoError(t, err)

	// Assert that the RunPythonScript function was called with the expected argument
	mockRunner.AssertExpectations(t)
}

// TestProcessData tests the ProcessData function
func TestProcessData(t *testing.T) {
	mockRunner := new(MockRunner)

	// Mock the behavior of the RunPythonScript function
	mockRunner.On("RunPythonScript", "scripts/process_data.py").Return(nil)

	// Call the function and assert it does not return an error
	err := ProcessData()
	assert.NoError(t, err)

	// Assert that the RunPythonScript function was called with the expected argument
	mockRunner.AssertExpectations(t)
}
