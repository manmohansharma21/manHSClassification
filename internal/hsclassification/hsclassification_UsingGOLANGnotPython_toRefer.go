package hsclassification

import (
	"fmt"
	"log"

	"github.com/manmohansharma21/hsclassifier/internal/categorizer"
	"github.com/xuri/excelize/v2"
)

// This function is not called from anywhere as it is not being used currently
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
