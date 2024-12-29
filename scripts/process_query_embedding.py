import pandas as pd
import numpy as np
from sklearn.metrics.pairwise import cosine_similarity
from sentence_transformers import SentenceTransformer
import json
import sys

# Debugging line to check passed arguments
print("Arguments passed to the script:", sys.argv)  # Debugging line

def load_master_embeddings():
    """
    Load the master embeddings from the CSV file and convert them to a numpy array.
    """
    try:
        # Load master embeddings CSV
        master_df = pd.read_csv('data/master_embeddings.csv')
        print("Loaded master_embeddings.csv successfully.")
        
        # Ensure that the embeddings are in numeric format (they should be lists of numbers)
        master_embeddings = master_df['embeddings'].apply(json.loads).apply(np.array)
        
        # Convert the list of arrays to a 2D numpy array
        master_embeddings = np.vstack(master_embeddings.values)
        print(f"Master embeddings loaded: {master_embeddings.shape}")
        
        return master_embeddings, master_df
    except Exception as e:
        print(f"Error loading master_embeddings.csv: {e}")
        return None, None

def generate_query_embedding(query):
    """
    Generate embedding for the input query using SentenceTransformer.
    """
    model = SentenceTransformer('all-MiniLM-L6-v2')
    query_embedding = model.encode([query])
    print(f"Query embedding generated for: {query}")
    return query_embedding

def process_queries(queries):
    """
    Process multiple queries and compute cosine similarity with master embeddings.
    """
    master_embeddings, master_df = load_master_embeddings()
    
    if master_embeddings is None or master_df is None:
        return

    # Process each query
    for query in queries:
        print(f"Processing query: {query}")
        
        # Generate the query embedding
        query_embedding = generate_query_embedding(query)
        
        # Compute cosine similarity between the query embedding and the master embeddings
        similarities = cosine_similarity(query_embedding, master_embeddings)
        
        # Find the best match using the highest similarity scor
        # Define a threshold for a close match
        threshold = 0.7
        
        # Get the top 5 most similar items (if any) based on similarity score
        top_indices = similarities.argsort()[0][-5:][::-1]
        
        found_match = False
        for idx in top_indices:
            if similarities[0][idx] >= threshold:
                best_match = master_df.iloc[idx]
                # print(f"Query: {query}")
                # print(f"Matched Product: {best_match['product_name']}")
                # print(f"HS Code: {best_match['hscode']}")
                # print(f"Description: {best_match['description']}")
                # print(f"Similarity: {similarities[0][idx]}")
        
                print(f"Best match: {best_match['product_name']}")
                print(f"HSCode: {best_match['hscode']}")
                found_match = True
                break  # Stop after printing the first (best) match
        
        if not found_match:
            print(f"No close match found for query: {query}. Please refine your query.")
            print(f"Query product name: {query}")  # Print the product name for which no match was found.
        print()  # Add a blank line between queries for better readability

if __name__ == "__main__":
    # List of queries
    # queries = [
    #     "laptop cover", 
    #     "smartphone case", 
    #     "tablet stand", 
    #     "laptop toy", 
    #     "laptop pictures", 
    #     "laptop machine", 
    #     "laptop sleeve", 
    #     "laptop charger"
    # ]
    
    # Get queries passed as command-line arguments
    if len(sys.argv) > 1:
        queries = sys.argv[1].split(",")
        # Process the queries
        process_queries(queries)
    else:
        print("No queries provided.")
