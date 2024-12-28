# python -m unittest test_convert_to_csv.py
# Importing necessary modules for testing
import unittest
from unittest.mock import patch, MagicMock
import pandas as pd
from io import StringIO

# Assuming the function to be tested is in a module named 'convert_to_csv_module'
# Uncomment the next line if the function is in a separate module
# from convert_to_csv_module import convert_to_csv


class TestConvertToCsv(unittest.TestCase):
    """
    This test case tests the `convert_to_csv` function by mocking the external dependencies
    such as file reading and writing. The goal is to ensure the function behaves correctly
    without actually reading or writing files during the test.
    """

    @patch('pandas.read_excel')  # Mocking the pandas `read_excel` function to avoid actual file reading
    @patch('pandas.DataFrame.to_csv')  # Mocking the `to_csv` method to prevent actual file writing
    @patch('builtins.print')  # Mocking the print function to capture the print output
    def test_convert_to_csv(self, mock_print, mock_to_csv, mock_read_excel):
        """
        This method tests the `convert_to_csv` function.
        It mocks the file reading and writing processes and verifies that the correct 
        methods are called with the expected arguments and outputs.
        """
        
        # Mock data for the Excel file, simulating the content that would be read from it
        mock_df = pd.DataFrame({
            'Description': ['item1', 'item2'],
            'Price': [100, 200]
        })
        
        # Setting the mock return value for `read_excel` to return the mock data frame
        mock_read_excel.return_value = mock_df

        # Call the function that we are testing
        convert_to_csv()

        # Verify that `read_excel` was called with the correct file path
        mock_read_excel.assert_called_once_with("data/descriptions.xlsx")

        # Verify that `to_csv` was called with the correct file path and parameters
        mock_to_csv.assert_called_once_with("data/descriptions.csv", index=False)

        # Verify that the expected output was printed
        mock_print.assert_any_call("File converted to descriptions.csv")
        mock_print.assert_any_call(mock_df)

# Run the tests when the script is executed
if __name__ == '__main__':
    unittest.main()
