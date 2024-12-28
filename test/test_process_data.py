# python -m unittest test_process_data.py

import unittest
from unittest.mock import patch, MagicMock
import pandas as pd
import os

# Assuming your function is in a module named 'process_data_module'
# from process_data_module import main, categorize_description


class TestProcessData(unittest.TestCase):
    """
    This test case tests the `process_data.py` script.
    It mocks the external dependencies such as file reading and writing,
    and ensures that the categorization logic works as expected.
    """

    @patch('pandas.read_excel')  # Mocking pandas read_excel to avoid actual file reading
    @patch('pandas.DataFrame.to_excel')  # Mocking the to_excel method to avoid actual file writing
    @patch('builtins.print')  # Mocking the print function to capture printed output
    @patch('os.path.exists')  # Mocking os.path.exists to simulate file existence
    def test_process_data(self, mock_exists, mock_print, mock_to_excel, mock_read_excel):
        """
        This method tests the `process_data.py` script.
        It mocks the file reading, file writing, and file existence check,
        and verifies that the correct methods are called with the expected arguments and outputs.
        """

        # Mocking os.path.exists to simulate that the input file exists
        mock_exists.return_value = True

        # Mock data for the Excel file, simulating the content that would be read from it
        mock_df = pd.DataFrame({
            'Description': ['laptop', 'laptop toy', 'unknown item']
        })
        mock_read_excel.return_value = mock_df

        # Call the main function (which is the entry point of the script)
        main()

        # Verify that `read_excel` was called with the correct file path
        mock_read_excel.assert_called_once_with("data/descriptions.xlsx")

        # Verify that `to_excel` was called with the correct file path and parameters
        mock_to_excel.assert_called_once_with("data/processed_descriptions.xlsx", index=False)

        # Verify that the correct print output was generated
        mock_print.assert_any_call("Processed data saved to data/processed_descriptions.xlsx")

        # Verify that the categorization was done correctly
        # Ensure that the correct categories were added to the DataFrame
        self.assertEqual(mock_df['Category'].iloc[0], 'laptop')
        self.assertEqual(mock_df['HS Code'].iloc[0], 'HS1234')
        self.assertEqual(mock_df['Category'].iloc[1], 'laptop toy')
        self.assertEqual(mock_df['HS Code'].iloc[1], 'HS5678')
        self.assertEqual(mock_df['Category'].iloc[2], 'Uncategorized')
        self.assertEqual(mock_df['HS Code'].iloc[2], 'HS0000')

    @patch('os.path.exists')  # Mocking os.path.exists to simulate file existence
    @patch('builtins.print')  # Mocking print to capture the output when file does not exist
    def test_file_not_found(self, mock_print, mock_exists):
        """
        This method tests the case where the input file does not exist.
        It ensures that the script prints the correct error message.
        """

        # Simulate that the input file does not exist
        mock_exists.return_value = False

        # Call the main function (which is the entry point of the script)
        main()

        # Verify that the correct print output was generated when the file is not found
        mock_print.assert_any_call("Input file data/descriptions.xlsx does not exist.")

# Run the tests when the script is executed
if __name__ == '__main__':
    unittest.main()
