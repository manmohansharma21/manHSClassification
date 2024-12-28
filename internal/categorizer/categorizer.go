package categorizer

import (
	"regexp"
	"strings"
)

// Category mapping with HS Codes
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
