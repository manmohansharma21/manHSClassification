import unittest
import os
import subprocess

class TestIntegration(unittest.TestCase):

    def test_integration(self):
        # Step 1: Generate sample data using the Python script
        result = subprocess.run(['python', 'scripts/create_sample_data.py'], capture_output=True, text=True)
        self.assertEqual(result.returncode, 0, f"Error: {result.stderr}")

        # Step 2: Process the data using the Python script
        result = subprocess.run(['python', 'scripts/process_data.py'], capture_output=True, text=True)
        self.assertEqual(result.returncode, 0, f"Error: {result.stderr}")

        # Step 3: Verify that the output file is created
        self.assertTrue(os.path.exists('data/processed_descriptions.xlsx'))

if __name__ == '__main__':
    unittest.main()
