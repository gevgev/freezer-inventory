Here is the final and complete version of the Backend PRD (v1.6) incorporating all changes since v1.1, along with the requested additions for Tagging APIs, Category Assignment APIs, and Category & Tag Retrieval APIs.

Backend PRD – Freezer Inventory Management API

Project Name: Freezer Inventory Management Backend
Date: February 13, 2025
Version: 1.6

1. Overview

The backend is a Golang-based API service for managing freezer inventory. It provides structured inventory tracking, role-based authentication, and AI-powered enhancements. Key components include:
	•	Item Definition – Represents a unique food item with general properties.
	•	Inventory Log – Tracks additions and removals of inventory.
	•	Categories & Tags – Organizes items for better filtering.
	•	Item-Category & Item-Tag Relationships – Supports multiple category and tag assignments.
	•	Inventory Reporting – Provides aggregated inventory calculations per item.
	•	Authentication & Authorization – Role-based access (admin vs. user).
	•	Cursor AI Integration – AI-assisted barcode lookup and image recognition.

The backend exposes a RESTful API for the iPhone app and Admin Portal, secured via JWT-based authentication.

2. Data Model

2.1 User Model (users table)

Stores registered users for authentication and authorization.

Field Name	Type	Description
id	UUID	Unique identifier for the user.
email	String	User's email (unique).
password_hash	String	Hashed password.
role	Enum	User role (admin, user).
created_at	Timestamp	Account creation timestamp.

2.2 Item Definition (items table)

Represents a unique type of food item.

Field Name	Type	Description
id	UUID	Unique identifier for the item.
name	String	Item name.
description	String	Optional item description.
barcode	String	Barcode (if applicable).
image_url	String	Image of the item.
packaging	String	Packaging type (box, bag, loose, etc.), if applicable.
created_at	Timestamp	Timestamp of item creation.
updated_at	Timestamp	Last update timestamp.
expiration_date	DATE	Expiration date of the item.
weight_unit	ENUM('g', 'kg', 'oz', 'lb')	Weight unit of the item.
deleted_at	TIMESTAMP	Timestamp of item deletion.

2.3 Inventory Log (inventory_log table)

Tracks all inventory changes.

Field Name	Type	Description
id	UUID	Unique identifier for the log entry.
item_id	UUID	References the items.id field.
timestamp	Timestamp	Date and time of the operation.
change	Integer	Signed numeric value (+ for additions, - for removals).
weight	Float	Optional weight information.
notes	String	Optional field for extra details.
weight_unit	ENUM('g', 'kg', 'oz', 'lb')	Weight unit of the inventory change.

2.4 Categories (categories table)

Defines categories that items can belong to.

Field Name	Type	Description
id	UUID	Unique identifier for the category.
name	String	Category name (e.g., "Meat", "Vegetables").
description	String	Optional category description.

2.5 Tags (tags table)

Defines available tags for items.

Field Name	Type	Description
id	UUID	Unique identifier for the tag.
name	String	Tag name (e.g., "Organic", "Frozen Meal").

2.6 Item-Category Relationship (item_categories table)

Maps items to one or more categories.

Field Name	Type	Description
id	UUID	Unique identifier.
item_id	UUID	References items.id.
category_id	UUID	References categories.id.

2.7 Item-Tag Relationship (item_tags table)

Maps items to multiple tags.

Field Name	Type	Description
id	UUID	Unique identifier.
item_id	UUID	References items.id.
tag_id	UUID	References tags.id.

3. API Endpoints

3.1 Category APIs

Method	Endpoint	Description
POST	/categories	Create a new category (admin only).
GET	/categories	Retrieve all categories.

3.2 Item-Category Management APIs

Method	Endpoint	Description
POST	/items/{id}/categories	Assign an item to one or more categories.
DELETE	/items/{id}/categories/{category_id}	Remove an item from a category.
GET	/items/categories/{category_id}	Retrieve all items in a category.
GET	/items/categories?ids=1,2,3	Retrieve items in multiple categories.

3.3 Tag APIs

Method	Endpoint	Description
POST	/tags	Create a new tag (admin only).
GET	/tags	Retrieve all available tags.

3.4 Item-Tag Management APIs

Method	Endpoint	Description
POST	/items/{id}/tags	Assign one or more tags to an item.
DELETE	/items/{id}/tags/{tag_id}	Remove a tag from an item.
GET	/items/tags/{tag_id}	Retrieve all items with a specific tag.
GET	/items/tags?ids=1,2,3	Retrieve items filtered by multiple tags.

4. Inventory Reporting

Formula:

Total Inventory for an Item = SUM(inventory_log.change) WHERE item_id = {ITEM_ID}

Example Response:

{
  "item_id": "123e4567-e89b-12d3-a456-426614174000",
  "item_name": "Chicken Breast",
  "total_quantity": 5,
  "total_weight": 10.5
}

5. API Request Examples

Adding Inventory

{
  "item_id": "123e4567-e89b-12d3-a456-426614174000",
  "change": 2,
  "quantity": 2,
  "weight": 1.5
}

Removing Inventory

{
  "item_id": "123e4567-e89b-12d3-a456-426614174000",
  "change": -1,
  "quantity": 1
}

6. Summary

This backend PRD fully integrates inventory reporting, category/tag relationships, API endpoints, and structured inventory tracking, providing real-time stock summaries and supporting role-based authentication.

GET /items - List all items
POST /items - Create new item
GET /items/{id} - Get item details
PUT /items/{id} - Update item
DELETE /items/{id} - Delete item

GET /inventory/{item_id}/status - Get current inventory
GET /inventory/{item_id}/history - Get inventory history

{
    "data": [
        {
            "id": "uuid",
            "name": "Chicken Breast",
            "expiration_date": "2025-03-01",
            "weight": 500,
            "weight_unit": "g"
        }
    ],
    "pagination": {
        "total": 100,
        "page": 1,
        "per_page": 20
    }
}

{
    "error": {
        "code": "INVALID_INPUT",
        "message": "Invalid item ID format",
        "details": {...}
    }
}

GET /items/search?q=chicken
GET /items?category=meat&expired=false

{
    "name": "required, string, max:255",
    "weight": "required, positive number",
    "expiration_date": "required, future date",
    "barcode": "optional, alphanumeric, max:50"
}

200 - Success
201 - Created
400 - Bad Request
401 - Unauthorized
403 - Forbidden
404 - Not Found
500 - Server Error

POST /auth/register
POST /auth/login
POST /auth/refresh
GET /users/me

7. Authentication

POST /auth/register
Request:
{
    "email": "user@example.com",
    "password": "min 8 chars, 1 uppercase, 1 number, 1 special char",
    "name": "string, required"
}
Response: 
{
    "token": "JWT token",
    "refresh_token": "refresh token",
    "user": { "id": "uuid", "email": "user@example.com", "role": "user" }
}

8. Rate Limiting
- API Rate Limit: 100 requests per minute per IP
- File Upload Rate Limit: 10 uploads per minute per user
Headers: 
  - X-RateLimit-Limit: 100
  - X-RateLimit-Remaining: 95
  - X-RateLimit-Reset: 1640995200

9. System Constraints
- Maximum items per user: 1000
- Maximum categories: 100
- Maximum tags per item: 20
- Image uploads: max 5MB, formats: JPG, PNG
- JWT token expiry: 1 hour
- Refresh token expiry: 30 days
