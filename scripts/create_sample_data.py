import pandas as pd

def create_sample_data():
    # Sample data
    data = {
        "Description": [
            "Laptop with high-speed processor",
            "A toy that looks like a laptop",
            "Photo of a vintage laptop",
            "High-end gaming laptop",
            "Children's educational laptop toy",
            "Artistic photo of a laptop",
            "Basic laptop for office work",
            "Wooden laptop toy",
            "Abstract painting of a laptop",
            "Laptop with a detachable screen"
        ]
    }

    # Create DataFrame
    df = pd.DataFrame(data)

    # Save to Excel
    df.to_excel("data/descriptions.xlsx", index=False)
    print("Sample descriptions.xlsx created in 'data/' directory.")

if __name__ == "__main__":
    create_sample_data()
