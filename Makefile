.PHONY: all generate-data process-data convert-to-csv run-go

all: generate-data convert-to-csv process-data run-go

generate-data:
	@echo "Generating test data..."
	python3 scripts/create_sample_data.py

convert-to-csv:
	@echo "Converting .xlsx to .csv..."
	python3 scripts/convert_to_csv.py

process-data:
	@echo "Processing descriptions..."
	python3 scripts/process_data.py

run-go:
	@echo "Running Go application..."
	go run cmd/main.go
