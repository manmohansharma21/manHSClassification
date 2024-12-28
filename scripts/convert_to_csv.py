import pandas as pd

def convert_to_csv():
    # Read the Excel file
    df = pd.read_excel("data/descriptions.xlsx")


    # Save as CSV
    df.to_csv("data/descriptions.csv", index=False)
    print("File converted to descriptions.csv")
    
    # Print the content
print(df)

if __name__ == "__main__":
    convert_to_csv()
