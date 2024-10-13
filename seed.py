import requests
import json
from datetime import datetime

# Define the API endpoint
url = "http://localhost:8080/api/v0/products"

# Create 10 sample products
products = [
    {
        "farmer_id": "farmer123",
        "name": f"Product {i}",
        "description": f"This is the description for Product {i}",
        "price": round(10 + i * 1.5, 2),
        "quantity": 100 - i * 5,
        "status": "available" if i % 2 == 0 else "out of stock",
        "image_url": f"http://example.com/image{i}.jpg",
        "shipping": True if i % 2 == 0 else False,
        "created_at": datetime.now().isoformat(),
        "updated_at": datetime.now().isoformat(),
    }
    for i in range(1, 11)  # Generate products with IDs 1 to 10
]


# Function to post a product to the API
def post_product(product):
    headers = {"Content-Type": "application/json"}
    response = requests.post(url, data=json.dumps(product), headers=headers)
    if response.status_code == 201:
        print(f"Successfully added product: {product['name']}")
    else:
        print(
            f"Failed to add product: {product['name']}, Status Code: {response.status_code}, Response: {response.text}"
        )


# Post each product to the API
for product in products:
    post_product(product)
