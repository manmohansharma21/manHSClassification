.PHONY: all generate-data generate-master-embeddings generate-query-embedding run-go clean

# Default target: Automates the entire workflow
all: generate-data generate-master-embeddings generate-query-embedding run-go

# Step 1: Generate test data
generate-data:
	@echo "Step 1: Generating test data..."
	python3 scripts/create_sample_data.py
	@echo "Test data generated successfully. Check the 'data/' directory for 'descriptions.xlsx'."

# Step 2: Generate master embeddings
generate-master-embeddings:
	@echo "Step 2: Generating master embeddings..."
	python3 scripts/generate_master_embeddings.py
	@echo "Master embeddings generated successfully. Check the 'data/' directory for 'master_embeddings.csv'."

# Step 3: Generate query embedding (processing multiple queries)
#python3 scripts/generate_query_embedding.py --query "laptop cover"
generate-query-embedding:
	@echo "Step 3: Generating query embeddings for input queries..."
	python3 scripts/process_query_embedding.py
	@echo "Query embeddings generated successfully. Check the console output for results."

# Step 4: Run the Go application
run-go:
	@echo "Step 4: Running Go application..."
	go run cmd/main.go
	@echo "Go application executed successfully."

# Clean up generated files
clean:
	@echo "Cleaning up generated files..."
	rm -f data/*.csv data/*.json data/*.xlsx
	@echo "Clean-up complete. All generated files removed."
