import pandas as pd

# Dummy data
data = [
    {"product_name": "Laptop", "hscode": "84713010", "description": "High-end personal laptop"},
    {"product_name": "Laptop Cover", "hscode": "42021910", "description": "Protective laptop cover"},
    {"product_name": "Laptop Toy", "hscode": "95030030", "description": "Toy laptop for kids"},
    {"product_name": "Laptop Picture", "hscode": "49119990", "description": "Printed picture of a laptop"},
    {"product_name": "Gaming Laptop", "hscode": "84713020", "description": "High-performance gaming laptop"},
    {"product_name": "Laptop Stand", "hscode": "39269099", "description": "Stand for holding laptops"},
    {"product_name": "Laptop Charger", "hscode": "85044010", "description": "Laptop charging device"},
    {"product_name": "Laptop Sleeve", "hscode": "42021210", "description": "Soft sleeve for laptops"},
    {"product_name": "Laptop Backpack", "hscode": "42029210", "description": "Backpack with laptop compartment"},
    {"product_name": "Laptop Keyboard", "hscode": "84716060", "description": "External keyboard for laptops"},
]

# Save to CSV
df = pd.DataFrame(data)
df.to_csv("data/master_table.csv", index=False)
