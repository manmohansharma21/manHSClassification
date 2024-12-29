package hsclassification

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

// LoadEmbeddings loads embeddings from a CSV file into a matrix
func LoadEmbeddings(filePath string) (*mat.Dense, []string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	// Extract embeddings and product names
	numRows := len(rows)
	numCols := len(rows[0]) - 1
	data := make([]float64, numRows*numCols)
	names := make([]string, numRows)

	for i, row := range rows {
		names[i] = row[0] // Assuming the first column contains product names
		for j := 1; j <= numCols; j++ {
			val, _ := strconv.ParseFloat(row[j], 64)
			data[i*numCols+j-1] = val
		}
	}

	// Create a dense matrix
	embeddings := mat.NewDense(numRows, numCols, data)
	return embeddings, names, nil
}

// FindClosestProduct performs similarity search on embeddings
func FindClosestProduct(queryEmbedding []float64, embeddings *mat.Dense, names []string) string {
	numRows, _ := embeddings.Dims()
	closestIndex := -1
	closestDistance := float64(1e9) // Initialize with a large number

	for i := 0; i < numRows; i++ {
		row := embeddings.RowView(i)
		distance := mat.Dot(mat.NewVecDense(len(queryEmbedding), queryEmbedding), row)
		if distance < closestDistance {
			closestDistance = distance
			closestIndex = i
		}
	}

	return names[closestIndex]
}

func main() {
	// Load embeddings
	embeddings, names, err := LoadEmbeddings("data/embeddings.csv")
	if err != nil {
		log.Fatalf("Failed to load embeddings: %v", err)
	}

	// Example query embedding
	queryEmbedding := []float64{0.1, 0.2, 0.3} // Replace with actual query embedding
	closestProduct := FindClosestProduct(queryEmbedding, embeddings, names)

	fmt.Printf("Closest product: %s\n", closestProduct)
}
