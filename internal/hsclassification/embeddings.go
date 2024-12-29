package hsclassification

import (
	"math"
	"strings"
)

// Generate a simple vector for a query (mock for demonstration)
func GenerateVector(input string) map[string]int {
	tokens := strings.Fields(strings.ToLower(input))
	vector := make(map[string]int)
	for _, token := range tokens {
		vector[token]++
	}
	return vector
}

// Calculate cosine similarity
func CosineSimilarity(vec1, vec2 map[string]int) float64 {
	dotProduct := 0
	mag1, mag2 := 0, 0

	for word, count1 := range vec1 {
		count2 := vec2[word]
		dotProduct += count1 * count2
		mag1 += count1 * count1
	}

	for _, count2 := range vec2 {
		mag2 += count2 * count2
	}

	if mag1 == 0 || mag2 == 0 {
		return 0
	}
	return float64(dotProduct) / (math.Sqrt(float64(mag1)) * math.Sqrt(float64(mag2)))
}

// Find the nearest match
func FindNearest(queryVec map[string]int, products []string) (string, float64) {
	bestMatch := ""
	bestScore := 0.0

	for _, product := range products {
		productVec := GenerateVector(product)
		score := CosineSimilarity(queryVec, productVec)
		if score > bestScore {
			bestMatch = product
			bestScore = score
		}
	}

	return bestMatch, bestScore
}
