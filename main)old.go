package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Category and HS Code mapping
var categories = map[string]string{
	"laptop":       "HS1234",
	"laptop toy":   "HS5678",
	"laptop photo": "HS9101",
}

// CategorizeDescription categorizes a description based on predefined categories
func CategorizeDescription(description string) (string, string) {
	description = strings.ToLower(description) // Case-insensitive matching
	for category, hscode := range categories {
		matched, _ := regexp.MatchString(`\b`+regexp.QuoteMeta(category)+`\b`, description)
		if matched {
			return category, hscode
		}
	}
	return "Uncategorized", "HS0000"
}

func main() {
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
			category, hscode := CategorizeDescription(description)
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
