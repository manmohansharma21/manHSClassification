// internal/hsclassification/hsclassification.go
package hsclassification

import (
	"fmt"
	"log"

	"github.com/manmohansharma21/hsclassifier/internal/runner"
)

func GenerateTestData() error {
	return runner.RunPythonScript("scripts/create_sample_data.py")
}

func ConvertToCSV() error {
	return runner.RunPythonScript("scripts/convert_to_csv.py")
}

func ProcessData() error {
	return runner.RunPythonScript("scripts/process_data.py")
}

func RunHSCodeClassification() {
	fmt.Println("Starting HS Code Mapper...")

	// Step 1: Generate test data
	if err := GenerateTestData(); err != nil {
		log.Fatalf("Error generating test data: %v", err)
	}

	// Step 2: Convert .xlsx to .csv
	if err := ConvertToCSV(); err != nil {
		log.Fatalf("Error converting .xlsx to .csv: %v", err)
	}

	// Step 3: Process the data
	if err := ProcessData(); err != nil {
		log.Fatalf("Error processing descriptions: %v", err)
	}

	fmt.Println("HS Code classification process completed successfully.")
}
