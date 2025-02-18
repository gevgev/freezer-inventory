# Freezer Inventory API Documentation

## Overview
RESTful API for managing freezer inventory, items, categories, and users.

## Base URL
```
http://localhost:8080
```

## Authentication
All endpoints except `/auth/*` require a JWT token:
```
Authorization: Bearer <token>
```

## API Endpoints

### 1. Authentication

#### Login
```http
POST /auth/login
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "Password123!"
}

Response: 200 OK
{
    "token": "eyJhbGc..."
}
```

#### Register
```http
POST /auth/register
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "Password123!"
}

Response: 201 Created
{
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "email": "user@example.com",
    "role": "user"
}
```

### 2. Items

#### List Items
```http
GET /api/items

Response: 200 OK
[
    {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "Chicken Breast",
        "description": "Organic chicken breast",
        "packaging": "Package",
        "weight_unit": "kg",
        "expiration_date": "2024-12-31",
        "categories": ["Meat"],
        "tags": ["Organic"]
    }
]
```

#### Get Single Item
```http
GET /api/items/:id

Response: 200 OK
{
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Chicken Breast",
    "description": "Organic chicken breast",
    "packaging": "Package",
    "weight_unit": "kg",
    "expiration_date": "2024-12-31",
    "categories": ["Meat"],
    "tags": ["Organic"]
}
```

#### Create Item
```http
POST /api/items
Content-Type: application/json

{
    "name": "Chicken Breast",
    "description": "Organic chicken breast",
    "packaging": "Package",
    "weight_unit": "kg",
    "expiration_date": "2024-12-31"
}

Response: 201 Created
```

### 3. Inventory Management

#### Get Item Status
```http
GET /api/inventory/:item_id/status

Response: 200 OK
{
    "item_id": "123e4567-e89b-12d3-a456-426614174000",
    "current_stock": 5,
    "total_weight": 2.5,
    "weight_unit": "kg"
}
```

#### Add Inventory Entry
```http
POST /api/inventory
Content-Type: application/json

{
    "item_id": "123e4567-e89b-12d3-a456-426614174000",
    "change": 1,
    "weight": 2.5,
    "weight_unit": "kg",
    "notes": "Added stock"
}

Response: 201 Created
```

### 4. Categories

#### List Categories
```http
GET /api/categories

Response: 200 OK
[
    {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "Meat",
        "description": "All types of meat products"
    }
]
```

#### Create Category (Admin Only)
```http
POST /api/categories
Content-Type: application/json

{
    "name": "Meat",
    "description": "All types of meat products"
}

Response: 201 Created
```

#### Update Category (Admin Only)
```http
PUT /api/categories/:id
Content-Type: application/json

{
    "name": "Updated Category",
    "description": "Updated description"
}

Response: 200 OK
```

### 5. Tags

#### List Tags
```http
GET /api/tags

Response: 200 OK
[
    {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "Organic"
    }
]
```

#### Create Tag (Admin Only)
```http
POST /api/tags
Content-Type: application/json

{
    "name": "Organic"
}

Response: 201 Created
```

#### Update Tag (Admin Only)
```http
PUT /api/tags/:id
Content-Type: application/json

{
    "name": "Updated Tag"
}

Response: 200 OK
```

### 6. User Management (Admin Only)

#### List Users
```http
GET /api/users

Response: 200 OK
[
    {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "email": "user@example.com",
        "role": "user"
    }
]
```

#### Create User
```http
POST /api/users
Content-Type: application/json

{
    "email": "newuser@example.com",
    "password": "Password123!",
    "role": "user"
}

Response: 201 Created
```

## Data Types

### Weight Units
```
Enum: ["kg", "g", "lb", "oz"]
```

### User Roles
```
Enum: ["admin", "user"]
```

## Error Responses

### Standard Error Format
```json
{
    "error": "Error message description"
}
```

### Common Status Codes
- `400` - Bad Request (Invalid input)
- `401` - Unauthorized (Invalid/missing token)
- `403` - Forbidden (Insufficient permissions)
- `404` - Not Found
- `500` - Internal Server Error

## Rate Limiting
API requests are limited to 100 requests per minute per IP address.

## Notes
- All timestamps are in ISO 8601 format
- All IDs are UUIDs
- Dates should be in YYYY-MM-DD format
- Admin endpoints require a token with admin role