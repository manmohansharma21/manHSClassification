# API Documentation

Currently, the HSClassifier project does not expose any RESTful APIs. However, it provides functionality via Go and Python scripts that can be run locally to process data.

## Available Functions

### `categorize_description(description: str) -> Tuple[str, str]`

This function takes a description string and categorizes it based on predefined categories. It returns the category and corresponding HS Code.

#### Parameters:

- `description` (str): The product description to be categorized.

#### Returns:

- Tuple[str, str]: A tuple containing the category and HS Code.

#### Example:

```python
category, hscode = categorize_description("Laptop with high-speed processor")
print(category)  # "laptop"
print(hscode)    # "HS1234"
```


### Python Scripts

The following Python scripts provide different functionalities:

#### `create_sample_data.py`

Generates a sample Excel file (`descriptions.xlsx`) containing product descriptions for testing.

#### `process_data.py`

Reads the** **`descriptions.xlsx` file, categorizes the descriptions using** **`categorize_description`, and saves the results to a new Excel file (`processed_descriptions.xlsx`).

#### `convert_to_csv.py`

Converts the processed Excel data into a CSV format.

### Running the Python Scripts

To execute the Python scripts:

<pre class="!overflow-visible"><div class="contain-inline-size rounded-md border-[0.5px] border-token-border-medium relative bg-token-sidebar-surface-primary dark:bg-gray-950"><div class="flex items-center text-token-text-secondary px-4 py-2 text-xs font-sans justify-between rounded-t-md h-9 bg-token-sidebar-surface-primary dark:bg-token-main-surface-secondary select-none">bash</div><div class="sticky top-9 md:top-[5.75rem]"><div class="absolute bottom-0 right-2 flex h-9 items-center"><div class="flex items-center rounded bg-token-sidebar-surface-primary px-2 font-sans text-xs text-token-text-secondary dark:bg-token-main-surface-secondary"><span class="" data-state="closed"><button class="flex gap-1 items-center select-none py-1" aria-label="Copy"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-sm"><path fill-rule="evenodd" clip-rule="evenodd" d="M7 5C7 3.34315 8.34315 2 10 2H19C20.6569 2 22 3.34315 22 5V14C22 15.6569 20.6569 17 19 17H17V19C17 20.6569 15.6569 22 14 22H5C3.34315 22 2 20.6569 2 19V10C2 8.34315 3.34315 7 5 7H7V5ZM9 7H14C15.6569 7 17 8.34315 17 10V15H19C19.5523 15 20 14.5523 20 14V5C20 4.44772 19.5523 4 19 4H10C9.44772 4 9 4.44772 9 5V7ZM5 9C4.44772 9 4 9.44772 4 10V19C4 19.5523 4.44772 20 5 20H14C14.5523 20 15 19.5523 15 19V10C15 9.44772 14.5523 9 14 9H5Z" fill="currentColor"></path></svg>Copy code</button></span></div></div></div><div class="overflow-y-auto p-4" dir="ltr"><code class="!whitespace-pre hljs language-bash">python scripts/create_sample_data.py
python scripts/process_data.py
python scripts/convert_to_csv.py</code></div></div></pre>
