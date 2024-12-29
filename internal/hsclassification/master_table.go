package hsclassification

var masterTable = map[string]string{
	"Laptop Machine": "84713010",
	"Laptop Picture": "49019900",
	"Laptop Toy":     "95030010",
	"Laptop Cover":   "42021250",
}

// Get product descriptions
func GetProductDescriptions() []string {
	descriptions := []string{}
	for product := range masterTable {
		descriptions = append(descriptions, product)
	}
	return descriptions
}

// Get HSCode for a given product description
func GetHSCodeForDescription(description string) string {
	return masterTable[description]
}
