#!/bin/bash

echo "Starting HS Code Mapper..."

# Step 1: Generate test data
echo "Generating test data..."
python3 scripts/create_sample_data.py
if [ $? -ne 0 ]; then
    echo "Error generating test data."
    exit 1
fi

# Step 2: Convert .xlsx to .csv
echo "Converting .xlsx to .csv..."
python3 scripts/convert_to_csv.py
if [ $? -ne 0 ]; then
    echo "Error converting .xlsx to .csv."
    exit 1
fi

# Step 3: Process the data
echo "Processing descriptions..."
python3 scripts/process_data.py
if [ $? -ne 0 ]; then
    echo "Error processing descriptions."
    exit 1
fi

# Step 4: Run the Go application
echo "Running Go application..."
go run cmd/main.go
if [ $? -ne 0 ]; then
    echo "Error running Go application."
    exit 1
fi

echo "HS Code Mapping process completed successfully."


: ` chmod +x run_project.sh
./run_project.sh
`

: '
 chmod +x run_project.sh
./run_project.sh

This is a multiline comment
in a shell script.
It will be ignored by the shell.
'
