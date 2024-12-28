import pandas as pd
import os

# Define categories and HS Codes
categories = {
    "laptop": "HS1234",
    "laptop toy": "HS5678",
    "laptop photo": "HS9101",
}

def categorize_description(description):
    """Categorize description based on predefined categories."""
    description = description.lower()
    for category, hscode in categories.items():
        if category in description:
            return category, hscode
    return "Uncategorized", "HS0000"

def main():
    # Paths
    input_file = "data/descriptions.xlsx"
    output_file = "data/processed_descriptions.xlsx"

    # Read Excel file
    if not os.path.exists(input_file):
        print(f"Input file {input_file} does not exist.")
        return

    df = pd.read_excel(input_file)

    # Categorize descriptions
    df[['Category', 'HS Code']] = df['Description'].apply(
        lambda x: pd.Series(categorize_description(x))
    )

    # Save to new Excel file
    df.to_excel(output_file, index=False)
    print(f"Processed data saved to {output_file}")

if __name__ == "__main__":
    main()
