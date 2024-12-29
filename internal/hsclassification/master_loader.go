package hsclassification

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

// LoadMasterEmbeddings loads the precomputed embeddings from the embeddings.csv file.
func LoadMasterEmbeddings(filepath string) (*mat.Dense, []string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	numRows := len(records)
	if numRows == 0 {
		return nil, nil, nil
	}
	numCols := len(records[0]) - 1 // Assuming the last column is the product name.

	embeddings := mat.NewDense(numRows, numCols, nil)
	productNames := make([]string, numRows)

	for i, record := range records {
		productNames[i] = record[numCols]
		for j := 0; j < numCols; j++ {
			value, err := strconv.ParseFloat(record[j], 64)
			if err != nil {
				log.Printf("Error parsing embedding value at row %d, col %d: %v", i, j, err)
				continue
			}
			embeddings.Set(i, j, value)
		}
	}

	return embeddings, productNames, nil
}
