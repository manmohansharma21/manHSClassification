from sentence_transformers import SentenceTransformer
import pandas as pd
import numpy as np
import argparse

# Load model
model = SentenceTransformer('all-MiniLM-L6-v2')

def generate_query_embedding(query):
    """
    Function to generate the embedding for a given query and return the embedding.
    """
    # Generate query embedding
    query_embedding = model.encode([query])
    return query_embedding

def main(query):
    """
    Main function to generate query embedding and print it.
    """
    # Generate query embedding
    query_embedding = generate_query_embedding(query)

    # Load the master table and its embeddings
    master_df = pd.read_csv("data/master_table.csv")
    master_embeddings = pd.read_csv("data/master_embeddings.csv").values

    # Find the closest match for the query using cosine similarity
    similarities = cosine_similarity(query_embedding, master_embeddings)
    best_match_idx = np.argmax(similarities)
    best_match = master_df.iloc[best_match_idx]

    # Print results
    print(f"Query: {query}")
    print(f"Matched Product: {best_match['product_name']}")
    print(f"HS Code: {best_match['hscode']}")
    print(f"Description: {best_match['description']}")
    print()

if __name__ == "__main__":
    # Parse query from command line arguments
    parser = argparse.ArgumentParser(description="Generate and process query embedding.")
    parser.add_argument('--query', type=str, required=True, help="Query to be processed.")
    args = parser.parse_args()

    # Call the main function with the query
    main(args.query)
