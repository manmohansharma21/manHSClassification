---

# HSClassifier Project

## Overview

The **HSClassifier** project processes product descriptions, categorizes them, and maps them to appropriate **HS codes**. It leverages both Go and Python scripts for different tasks: generating test data, processing Excel files, converting data formats, and categorizing descriptions. The goal is to demonstrate how Go and Python can work together in a data processing pipeline.

---

## Project Structure

The project follows a modular structure to separate concerns and improve maintainability:

```
hsclassifier/
├── cmd/
│   └── main.go                     # Entry point for the Go application
│   └── main_test.go                # Unit tests for main.go (in the same directory)
├── internal/
│   ├── categorizer/                # Categorization logic and tests
│   │   ├── categorizer.go
│   │   └── categorizer_test.go
│   └── runner/                     # Python script runner and tests
│       ├── runner.go
│       └── runner_test.go
├── scripts/                         # Python scripts for data handling
│   ├── create_sample_data.py
│   ├── convert_to_csv.py
│   └── process_data.py
├── data/                            # Sample input/output data files
│   └── descriptions.xlsx
├── test/                            # Unit tests for the Go and Python components
│   ├── test_convert_to_csv.py
│   ├── test_process_data.py
│   ├── test_create_sample_data.py
│   ├── test_main.go
│   └── integration_test.go
├── go.mod                           # Go module file
├── go.sum                           # Go module checksum file
└── README.md                        # Project documentation
```

## Components

### 1. **cmd/main.go**

The entry point for the Go application. It orchestrates the execution of the entire workflow, including running Python scripts and calling Go functions for categorization.

### 2. **internal/runner/runner.go**

This package contains the `RunPythonScript` function, which is responsible for executing Python scripts from Go. It ensures that Python scripts are executed in the correct order.

### 3. **internal/categorizer/categorizer.go**

This package contains the logic for categorizing descriptions. The categorization logic is implemented in Go, while Python scripts handle data processing tasks.

### 4. **scripts/**

This directory contains Python scripts for handling data processing tasks:

- `create_sample_data.py`: Generates sample data for testing.
- `process_data.py`: Processes data for categorization.
- `convert_to_csv.py`: Converts the processed data from Excel to CSV.

### 5. **data/**

Contains input and output files:

- `descriptions.xlsx`: The input Excel file generated by `create_sample_data.py`.
- `processed_descriptions.xlsx`: The output file after processing the data.

### 6. **run_project.sh**

A shell script that orchestrates the workflow by running the necessary Python scripts and the Go application in sequence.

---

## Setup Instructions

### Step 1: Install Dependencies

**Install Go dependencies:**

```bash
go mod tidy
```

**Install Python dependencies (if required for the Python scripts):**

```bash
pip install -r requirements.txt
```

### Step 2: Running the Project

1. **Generate test data**: Use the `create_sample_data.py` Python script to generate sample data.

   ```bash
   python scripts/create_sample_data.py
   ```
2. **Process the data**: Use the `process_data.py` script to process the data for categorization.

   ```bash
   python scripts/process_data.py
   ```
3. **Run the Go application**: This will execute the Go application, which orchestrates the execution of the Python scripts and categorizes the descriptions.

   ```bash
   go run cmd/main.go
   ```

---

## How to Run the Project

### For Go Application:

To run the Go application, use:

```bash
go run cmd/main.go
```

### For Python Scripts:

To run the individual Python scripts, use:

```bash
python scripts/create_sample_data.py
python scripts/convert_to_csv.py
python scripts/process_data.py
```

### For Tests:

**Go tests:**

```bash
go test ./...
```

**Python tests:**

```bash
python -m unittest discover -s test
```

### Step 5: Run the Integration Tests

After setting up the project, you can run the integration tests to verify the interaction between the Go application and Python scripts.

1. **For Go integration tests**:

   Run the Go integration tests using:

   ```bash
   go test -v test/integration_test.go
   ```
2. **For Python integration tests**:

   Run the Python integration tests using:

   ```bash
   python -m unittest test/integration_test.py
   ```

---

## Git Operations

1. **Add and commit all files**:

   ```bash
   git add .
   git commit -m "Initial project structure"
   ```
2. **Push to a remote repository (GitHub, GitLab, etc.)**:

   ```bash
   git remote add origin <your-repository-url>
   git push -u origin master
   ```

---

## Project Workflow

The overall workflow is as follows:

1. **Generate test data** using the `create_sample_data.py` Python script.
2. **Process the data** using the `process_data.py` script.
3. **Run the Go application** to use the processed data for categorization.

---

## Future Enhancements

- **Improve categorization logic** in Go.
- **Integrate additional Python scripts** for advanced data processing.
- **Enhance error handling** and logging for both Python and Go scripts.
- **Add more data formats** for input and output (e.g., JSON, XML).

---

## License

This project is licensed under the man_MIT License.

---

### References to Other Documentation Files

- [Installation Guide](docs/installation.md): Instructions for installing and setting up the project.
- [Usage Guide](docs/usage.md): Instructions on how to use the project, with examples.
- [API Documentation](docs/api-reference.md): API documentation if applicable.
- [Architecture Overview](docs/architecture.md): Technical architecture or design documents.

---


Note:

Orchestrate Meaning

In this context, orchestrate means to coordinate or arrange the execution of multiple tasks (like running Python scripts and Go programs) in a specific order. This ensures the dependencies between tasks are respected. For example:

First Task: Generate test data using the Python script.

Second Task: Process the data using another Python script.

Final Task: Run the Go application to use the processed data.
