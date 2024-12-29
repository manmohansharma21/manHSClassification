import pandas as pd
import numpy as np
import openai
import json
import sys
from sklearn.metrics.pairwise import cosine_similarity
from dotenv import load_dotenv
import os

# Load environment variables from .env file
load_dotenv()

# Set up OpenAI API key
# Get the API key from environment variable
openai.api_key = os.getenv("OPENAI_API_KEY")


def generate_query_embedding(query):
    """
    Generate embedding for the input query using OpenAI's new API.
    """
    try:
        # Call OpenAI API to generate embeddings using the new API
        response = openai.embeddings.create(
            model="text-embedding-ada-002",  # Use OpenAI's embeddings model
            input=query
        )
        query_embedding = response['data'][0]['embedding']
        print(f"Query embedding generated for: {query}")
        return np.array(query_embedding)  # Convert to numpy array for consistency
    except Exception as e:
        print(f"Error generating query embedding: {e}")
        return None

def load_master_embeddings():
    """
    Generate embeddings for the master data using OpenAI's new API.
    """
    try:
        # Load the master data (assuming you have product descriptions)
        master_df = pd.read_csv('data/master_table.csv')
        print("Loaded master_table.csv successfully.")
        
        # Generate embeddings for each product description using the new API
        master_embeddings = []
        for description in master_df['description']:
            response = openai.embeddings.create(
                model="text-embedding-ada-002",
                input=description
            )
            master_embeddings.append(response['data'][0]['embedding'])
            
        # Convert the list of embeddings to a 2D numpy array
        master_embeddings = np.array(master_embeddings)
        print(f"Master embeddings generated: {master_embeddings.shape}")
        
        return master_embeddings, master_df
    except Exception as e:
        print(f"Error generating master embeddings: {e}")
        return None, None

def process_queries(queries):
    """
    Process multiple queries and compute cosine similarity with master embeddings.
    """
    master_embeddings, master_df = load_master_embeddings()
    
    if master_embeddings is None or master_df is None:
        return

    # Process each query
    for query in queries:
        # Generate the query embedding
        query_embedding = generate_query_embedding(query)
        
        if query_embedding is None:
            print(f"Skipping query: {query} due to embedding generation failure.")
            continue
        
        # Compute cosine similarity between the query embedding and the master embeddings
        similarities = cosine_similarity(query_embedding.reshape(1, -1), master_embeddings)
        
        # Find the best match using the highest similarity score
        threshold = 0.7
        top_indices = similarities.argsort()[0][-5:][::-1]
        
        found_match = False
        for idx in top_indices:
            if similarities[0][idx] >= threshold:
                best_match = master_df.iloc[idx]
                print(f"Query: {query}")
                print(f"Matched Product: {best_match['product_name']}")
                print(f"HS Code: {best_match['hscode']}")
                print(f"Description: {best_match['description']}")
                print(f"Similarity: {similarities[0][idx]}")
                found_match = True
        
        if not found_match:
            print(f"No close match found for query: {query}. Please refine your query.")
        print()  # Add a blank line between queries for better readability

if __name__ == "__main__":
    # Get queries passed as command-line arguments
    queries = sys.argv[1].split(",")
    
    # Process the queries
    process_queries(queries)
