import pandas as pd
from sentence_transformers import SentenceTransformer

def detect_delimiter(file_path):
    """
    Function to automatically detect the delimiter in a CSV file.
    It checks for common delimiters (comma, tab, semicolon, space) and returns the correct one.
    """
    delimiters = [',', '\t', ';', ' ']
    for delimiter in delimiters:
        try:
            # Attempt to load the file with the given delimiter
            df = pd.read_csv(file_path, delimiter=delimiter)
            # If no null columns, it suggests correct delimiter
            if df.columns.isnull().sum() == 0:
                print(f"Delimiter detected: {repr(delimiter)}")
                return delimiter
        except Exception as e:
            continue  # If this delimiter fails, move to the next one
    
    raise ValueError(f"Could not detect delimiter for file: {file_path}")

def generate_master_embeddings():
    """
    Function to generate and save master embeddings from descriptions in the master table.
    """
    # Step 1: Load the descriptions from the Excel file
    try:
        df = pd.read_excel("data/descriptions.xlsx")
        print("Loaded descriptions.xlsx successfully.")
    except Exception as e:
        print(f"Error loading descriptions.xlsx: {e}")
        return

    # Normalize column names to lowercase for consistency
    df.columns = df.columns.str.strip().str.lower()

    # Print the column names to debug
    print("Columns in descriptions DataFrame:", df.columns)

    # Ensure 'description' column exists
    if 'description' not in df.columns:
        raise ValueError("The 'description' column is missing in the descriptions data.")

    # Step 2: Load the master table (detect and handle the correct delimiter)
    try:
        master_file_path = "data/master_table.csv"
        delimiter = detect_delimiter(master_file_path)
        master_df = pd.read_csv(master_file_path, delimiter=delimiter)
        print("Loaded master_table.csv successfully.")
    except Exception as e:
        print(f"Error loading master_table.csv: {e}")
        return

    # Normalize column names to lowercase for consistency
    master_df.columns = master_df.columns.str.strip().str.lower()

    # Print the column names of the master table for debugging
    print("Columns in master table DataFrame:", master_df.columns)

    # If columns are not split correctly (e.g., using tab delimiter), try reloading with tab delimiter explicitly
    if len(master_df.columns) == 1 and '\t' in master_df.columns[0]:
        print("Columns appear to be tab-separated, reloading with tab delimiter.")
        master_df = pd.read_csv(master_file_path, delimiter='\t')

    # Print the columns again after reloading with tab delimiter
    print("Columns after reloading with tab delimiter:", master_df.columns)

    # Ensure 'description' column is present in the master table
    if 'description' not in master_df.columns:
        raise ValueError("The 'description' column is missing in the master table.")

    # Step 3: Initialize the SentenceTransformer model
    model = SentenceTransformer('all-MiniLM-L6-v2')

    # Step 4: Generate embeddings for the 'description' column in the master table
    embeddings = model.encode(master_df['description'].tolist())
    print(f"Generated {len(embeddings)} embeddings for product descriptions.")

    # Step 5: Add the embeddings as a new column to the master table
    master_df['embeddings'] = embeddings.tolist()

    # Step 6: Save the embeddings to a new CSV file
    try:
        master_df.to_csv("data/master_embeddings.csv", index=False)
        print("Embeddings generated and saved to 'data/master_embeddings.csv'.")
    except Exception as e:
        print(f"Error saving master_embeddings.csv: {e}")

# Call the function to generate master embeddings
generate_master_embeddings()
