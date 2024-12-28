// go run cmd/main.go

package main

import (
	"fmt"
	"log"

	"github.com/manmohansharma21/hsclassifier/internal/categorizer"
	"github.com/manmohansharma21/hsclassifier/internal/runner"
	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("Starting HS Code Mapper...")

	// Step 1: Generate test data
	fmt.Println("Generating test data...")
	err := runner.RunPythonScript("scripts/create_sample_data.py")
	if err != nil {
		log.Fatalf("Error generating test data: %v", err)
	}

	// Step 2: Convert .xlsx to .csv
	fmt.Println("Converting .xlsx to .csv...")
	err = runner.RunPythonScript("scripts/convert_to_csv.py")
	if err != nil {
		log.Fatalf("Error converting .xlsx to .csv: %v", err)
	}

	// Step 3: Process the data
	fmt.Println("Processing descriptions...")
	err = runner.RunPythonScript("scripts/process_data.py")
	if err != nil {
		log.Fatalf("Error processing descriptions: %v", err)
	}

	// Step 4: Run the Go application
	fmt.Println("Running Go application...")
	// Call your Go functions for the next steps
	// e.g., categorization and mapping

	fmt.Println("HS Code classification process completed successfully.")
}

// This function is not called from anywehere as it is not being used currently
// Python Pandas being used
// to attain the same functionality.
func SameTaskUsingGolang_JustToReferLater() {
	// Open the input Excel file
	inputFile := "descriptions.xlsx"
	outputFile := "categorized_descriptions.xlsx"
	f, err := excelize.OpenFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}

	// Read rows from the first sheet
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to read rows: %v", err)
	}

	// Add headers for the new columns
	rows[0] = append(rows[0], "Category", "HS Code")

	// Process each description
	for i, row := range rows[1:] {
		if len(row) > 0 {
			description := row[0]
			category, hscode := categorizer.CategorizeDescription(description)
			rows[i+1] = append(row, category, hscode)
		}
	}

	// Create a new Excel file for output
	output := excelize.NewFile()
	outputSheet := "Sheet1"
	output.NewSheet(outputSheet)

	// Write rows to the new file
	for i, row := range rows {
		for j, colCell := range row {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+1)
			if err := output.SetCellValue(outputSheet, cell, colCell); err != nil {
				log.Fatalf("Failed to write cell: %v", err)
			}
		}
	}

	// Save the output file
	if err := output.SaveAs(outputFile); err != nil {
		log.Fatalf("Failed to save output file: %v", err)
	}

	fmt.Printf("Categorized data saved to %s\n", outputFile)
}


//
package runner

import (
    "fmt"
    "os/exec"
)

// RunPythonScript executes a Python script
func RunPythonScript(scriptPath string) error {
    cmd := exec.Command("python3", scriptPath)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error running script %s: %v\nOutput: %s", scriptPath, err, string(output))
    }
    fmt.Printf("Script %s executed successfully.\n", scriptPath)
    return nil
}
