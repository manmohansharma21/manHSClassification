# Architecture Overview

The HSClassifier project is built using Go for the core application logic and Python for data processing. The project involves reading data from Excel files, processing the data, categorizing it, and then converting it into different formats. The architecture follows a modular design where each component has a clear responsibility.

## Components

1. **Go Application (`cmd/main.go`)**

   - The Go application is the entry point for the project.
   - It orchestrates the execution of Python scripts and coordinates the entire workflow.
2. **Categorization Logic (`internal/categorizer/categorizer.go`)**

   - This module is responsible for categorizing descriptions based on predefined categories.
   - It defines the `categorize_description` function, which is used to categorize product descriptions.
3. **Python Scripts (`scripts/`)**

   - The Python scripts handle various tasks:
     - `create_sample_data.py`: Generates sample descriptions for testing.
     - `process_data.py`: Processes the data and categorizes it.
     - `convert_to_csv.py`: Converts the categorized data to CSV format.
4. **Data (`data/`)**

   - This directory contains the input and output data files:
     - `descriptions.xlsx`: The input file containing raw descriptions.
     - `processed_descriptions.xlsx`: The output file containing categorized descriptions.
5. **Tests (`test/`)**

   - The tests are organized into separate files for Go and Python:
     - Go tests are located in `test/test_main.go`.
     - Python tests are located in `test/test_convert_to_csv.py`, `test/test_process_data.py`, and `test/test_create_sample_data.py`.

## Workflow

1. **Step 1: Generate Sample Data**

   - The Python script `create_sample_data.py` generates sample descriptions and saves them to `data/descriptions.xlsx`.
2. **Step 2: Process Data**

   - The `process_data.py` script reads the sample data, categorizes it using the `categorize_description` function, and saves the results to `data/processed_descriptions.xlsx`.
3. **Step 3: Run the Go Application**

   - The Go application reads the processed data, uses it for categorization, and outputs the final results.

## Diagram

Below is a simplified diagram of the project's architecture:

+---------------------+ +------------------+ +------------------------+ | Go Application | -----> | Categorizer | -----> | Python Scripts | | (cmd/main.go) | | (categorizer.go) | | (create_sample_data.py,| | | | | | process_data.py, | | | | | | convert_to_csv.py) | +---------------------+ +------------------+ +------------------------+

HSClassifier/
├── Makefile                     # Automates the workflow
├── cmd/
│   └── main.go                  # Entry point for the Go application
├── internal/
│   ├── hsclassification/
│   │   ├── classification.go    # Main classification logic
│   │   ├── embeddings.go        # Vectorization and embedding logic
│   │   ├── master_loader.go     # Logic for loading master embeddings
│   │   ├── master_table.go      # Handles master table data
│   │   ├── query_handler.go     # Handles runtime query embedding
├── scripts/                     # Python scripts for data handling
│   ├── create_sample_data.py    # Generates test data
│   ├── generate_master_embeddings.py # Creates master embeddings
│   ├── generate_query_embedding.py  # Processes input query embeddings
│   ├── convert_to_csv.py        # Converts data to CSV format
│   ├── process_data.py          # Processes descriptions or other input data
├── data/                        # Sample input/output data files
│   ├── descriptions.xlsx        # Original descriptions
│   ├── master_table.csv         # Master table with processed data
│   ├── master_embeddings.csv    # Generated master embeddings
│   ├── query_embedding.json     # Generated query embedding
├── test/                        # Unit tests for Go and Python components
│   ├── test_create_sample_data.py
│   ├── test_convert_to_csv.py
│   ├── test_process_data.py
│   ├── main_test.go             # Unit tests for main.go
│   ├── integration_test.go      # Integration tests for end-to-end validation
├── go.mod                       # Go module file for dependency management
├── go.sum                       # Dependency checksums
└── README.md                    # Project documentation


---

These are the four documentation files (`installation.md`,** **`usage.md`,** **`api-reference.md`, and** **`architecture.md`) in a format ready for direct copying into your project. Each file contains the necessary content for your repository’s** **`docs/`folder.
