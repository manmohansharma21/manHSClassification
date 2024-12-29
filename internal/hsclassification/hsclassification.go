// internal/hsclassification/hsclassification.go
package hsclassification

import (
	"fmt"
	"log"

	// "github.com/manmohansharma21/hsclassifier/internal/hsclassification/embeddings"
	// "github.com/manmohansharma21/hsclassifier/internal/hsclassification/master_table"
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

func RunHSCodeClassification_older() {
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

func RunHSCodeClassification(query string) {
	// Load master table data
	products := GetProductDescriptions() //master_table.GetProductDescriptions()

	// Generate embeddings for the query
	queryVector := GenerateVector(query) //embeddings.GenerateVector(query)

	// Find the best match
	bestMatch, bestScore := FindNearest(queryVector, products) // embeddings.FindNearest(queryVector, products)

	if bestScore > 0.7 {
		fmt.Printf("Best match: %s\n", bestMatch)
		hsCode := GetHSCodeForDescription(bestMatch) //master_table.GetHSCodeForDescription(bestMatch)
		fmt.Printf("HSCode: %s\n", hsCode)
	} else {
		fmt.Println("No close match found. Please refine your query.")
	}
}
