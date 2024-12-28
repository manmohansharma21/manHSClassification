package categorizer

import (
	"testing"
)

func TestCategorizeDescription(t *testing.T) {
	tests := []struct {
		name             string
		description      string
		expectedCategory string
		expectedHSCode   string
	}{
		{
			name:             "Valid category 'laptop'",
			description:      "This is a laptop.",
			expectedCategory: "laptop",
			expectedHSCode:   "HS1234",
		},
		{
			name:             "Valid category 'laptop toy'",
			description:      "A laptop toy for kids.",
			expectedCategory: "laptop toy",
			expectedHSCode:   "HS5678",
		},
		{
			name:             "Valid category 'laptop photo'",
			description:      "Laptop photo for sale.",
			expectedCategory: "laptop photo",
			expectedHSCode:   "HS9101",
		},
		{
			name:             "No match category",
			description:      "This is a phone.",
			expectedCategory: "Uncategorized",
			expectedHSCode:   "HS0000",
		},
		{
			name:             "Mixed case match",
			description:      "A LAPTOP TOY for children.",
			expectedCategory: "laptop toy",
			expectedHSCode:   "HS5678",
		},
		{
			name:             "Partial match",
			description:      "This is a small laptop toy.",
			expectedCategory: "laptop toy",
			expectedHSCode:   "HS5678",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category, hscode := CategorizeDescription(tt.description)
			if category != tt.expectedCategory || hscode != tt.expectedHSCode {
				t.Errorf("CategorizeDescription() = (%v, %v), want (%v, %v)", category, hscode, tt.expectedCategory, tt.expectedHSCode)
			}
		})
	}
}
