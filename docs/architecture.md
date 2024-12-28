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

<pre class="!overflow-visible"><div class="contain-inline-size rounded-md border-[0.5px] border-token-border-medium relative bg-token-sidebar-surface-primary dark:bg-gray-950"><div class="flex items-center text-token-text-secondary px-4 py-2 text-xs font-sans justify-between rounded-t-md h-9 bg-token-sidebar-surface-primary dark:bg-token-main-surface-secondary select-none">markdown</div><div class="sticky top-9 md:top-[5.75rem]"><div class="absolute bottom-0 right-2 flex h-9 items-center"><div class="flex items-center rounded bg-token-sidebar-surface-primary px-2 font-sans text-xs text-token-text-secondary dark:bg-token-main-surface-secondary"><span class="" data-state="closed"><button class="flex gap-1 items-center select-none py-1" aria-label="Copy"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-sm"><path fill-rule="evenodd" clip-rule="evenodd" d="M7 5C7 3.34315 8.34315 2 10 2H19C20.6569 2 22 3.34315 22 5V14C22 15.6569 20.6569 17 19 17H17V19C17 20.6569 15.6569 22 14 22H5C3.34315 22 2 20.6569 2 19V10C2 8.34315 3.34315 7 5 7H7V5ZM9 7H14C15.6569 7 17 8.34315 17 10V15H19C19.5523 15 20 14.5523 20 14V5C20 4.44772 19.5523 4 19 4H10C9.44772 4 9 4.44772 9 5V7ZM5 9C4.44772 9 4 9.44772 4 10V19C4 19.5523 4.44772 20 5 20H14C14.5523 20 15 19.5523 15 19V10C15 9.44772 14.5523 9 14 9H5Z" fill="currentColor"></path></svg>Copy code</button></span></div></div></div><div class="overflow-y-auto p-4" dir="ltr"><code class="!whitespace-pre hljs language-markdown">
This modular architecture allows for easy updates and extension of functionality, such as adding new categorization logic or additional data processing steps.

## Future Enhancements

- **Improve categorization logic**: Add more categories or machine learning-based categorization.
- **Enhance error handling**: Implement better error reporting and handling in Python and Go scripts.
- **Integration with databases**: Store the processed data in a database for easier access and analysis.
</code></div></div></pre>

---

These are the four documentation files (`installation.md`,** **`usage.md`,** **`api-reference.md`, and** **`architecture.md`) in a format ready for direct copying into your project. Each file contains the necessary content for your repository’s** **`docs/`folder.
