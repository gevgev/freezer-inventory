# Freezer Inventory API Documentation

## Overview
RESTful API service for managing freezer inventory, providing endpoints for item tracking, inventory management, user administration, and categorization.

## Base URL
```
http://localhost:8080
```

## Authentication
All endpoints except `/auth/*` require JWT authentication via Bearer token:
```http
Authorization: Bearer <token>
```

## API Endpoints

### 1. Authentication

#### Login
```http
POST /auth/login
Content-Type: application/json

Request:
{
    "email": "user@example.com",
    "password": "Password123!"
}

Response: 200 OK
{
    "token": "eyJhbGc...",
    "user": {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "email": "user@example.com",
        "role": "user"
    }
}
```

#### Register
```http
POST /auth/register
Content-Type: application/json

Request:
{
    "email": "user@example.com",
    "password": "Password123!"
}

Response: 201 Created
{
    "token": "eyJhbGc...",
    "user": {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "email": "user@example.com",
        "role": "user"
    }
}
```

#### Refresh Token
```http
POST /auth/refresh
Authorization: Bearer <refresh_token>

Response: 200 OK
{
    "token": "eyJhbGc..."
}
```

### 2. Item Management

#### List Items
```http
GET /api/items
Authorization: Bearer <token>

Query Parameters:
- category_id (optional): Filter by category
- tag_id (optional): Filter by tag
- expired (optional): Filter by expiration status
- page (optional): Page number for pagination
- per_page (optional): Items per page

Response: 200 OK
{
    "data": [
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
    ],
    "pagination": {
        "total": 100,
        "page": 1,
        "per_page": 20
    }
}
```

#### Get Item
```http
GET /api/items/:id
Authorization: Bearer <token>

Response: 200 OK
{
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Chicken Breast",
    "description": "Organic chicken breast",
    "packaging": "Package",
    "weight_unit": "kg",
    "expiration_date": "2024-12-31",
    "categories": ["Meat"],
    "tags": ["Organic"],
    "current_stock": 5
}
```

#### Create Item
```http
POST /api/items
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "name": "Chicken Breast",
    "description": "Organic chicken breast",
    "packaging": "Package",
    "weight_unit": "kg",
    "expiration_date": "2024-12-31"
}

Response: 201 Created
{
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Chicken Breast",
    ...
}
```

#### Update Item
```http
PUT /api/items/:id
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "name": "Updated Chicken Breast",
    "description": "Updated description",
    "packaging": "Package",
    "weight_unit": "kg",
    "expiration_date": "2024-12-31"
}

Response: 200 OK
```

#### Delete Item
```http
DELETE /api/items/:id
Authorization: Bearer <token>

Response: 200 OK
{
    "message": "Item deleted successfully"
}
```

#### Search Items
```http
GET /api/items/search?q=chicken
Authorization: Bearer <token>

Response: 200 OK
{
    "data": [...],
    "pagination": {...}
}
```

### 3. Inventory Management

#### Get Item Status
```http
GET /api/inventory/:item_id/status
Authorization: Bearer <token>

Response: 200 OK
{
    "item_id": "123e4567-e89b-12d3-a456-426614174000",
    "current_stock": 5,
    "total_weight": 2.5,
    "weight_unit": "kg"
}
```

#### Get Item History
```http
GET /api/inventory/:item_id/history
Authorization: Bearer <token>

Response: 200 OK
[
    {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "timestamp": "2024-03-14T12:00:00Z",
        "change": 5,
        "weight": 2.5,
        "weight_unit": "kg",
        "notes": "Initial stock"
    }
]
```

#### Add Inventory Entry
```http
POST /api/inventory
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "item_id": "123e4567-e89b-12d3-a456-426614174000",
    "change": 1,
    "weight": 2.5,
    "weight_unit": "kg",
    "notes": "Added stock"
}

Response: 201 Created
```

### 4. Category Management

#### List Categories
```http
GET /api/categories
Authorization: Bearer <token>

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
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "name": "Meat",
    "description": "All types of meat products"
}

Response: 201 Created
```

#### Update Category (Admin Only)
```http
PUT /api/categories/:id
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "name": "Updated Category",
    "description": "Updated description"
}

Response: 200 OK
```

### 5. Tag Management

#### List Tags
```http
GET /api/tags
Authorization: Bearer <token>

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
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "name": "Organic"
}

Response: 201 Created
```

#### Update Tag (Admin Only)
```http
PUT /api/tags/:id
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "name": "Updated Tag"
}

Response: 200 OK
```

### 6. User Management (Admin Only)

#### List Users
```http
GET /api/users
Authorization: Bearer <token>

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
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "email": "newuser@example.com",
    "password": "Password123!",
    "role": "user"
}

Response: 201 Created
```

#### Update User
```http
PUT /api/users/:id
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "email": "updated@example.com",
    "password": "NewPassword123!",
    "role": "user"
}

Response: 200 OK
```

#### Delete User
```http
DELETE /api/users/:id
Authorization: Bearer <token>

Response: 200 OK
{
    "message": "User deleted successfully"
}
```

### Item Category Management

#### Add Categories to Item
```http
POST /api/items/:id/categories
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "category_ids": ["uuid1", "uuid2"]
}

Response: 200 OK
{
    "message": "Categories added successfully"
}
```

#### Remove Category from Item
```http
DELETE /api/items/:id/categories/:category_id
Authorization: Bearer <token>

Response: 200 OK
{
    "message": "Category removed successfully"
}
```

### Item Tag Management

#### Add Tags to Item
```http
POST /api/items/:id/tags
Authorization: Bearer <token>
Content-Type: application/json

Request:
{
    "tag_ids": ["uuid1", "uuid2"]
}

Response: 200 OK
{
    "message": "Tags added successfully"
}
```

#### Remove Tag from Item
```http
DELETE /api/items/:id/tags/:tag_id
Authorization: Bearer <token>

Response: 200 OK
{
    "message": "Tag removed successfully"
}
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

### Validation Rules
```
Email: required, valid email format
Password: min 8 chars, 1 uppercase, 1 number, 1 special char
Name: required, max 255 chars
Description: optional, max 1000 chars
```

## Error Responses

### Standard Error Format
```json
{
    "error": {
        "code": "INVALID_INPUT",
        "message": "Invalid item ID format",
        "details": {...}
    }
}
```

### Common Status Codes
- `200` - Success
- `201` - Created
- `400` - Bad Request (Invalid input)
- `401` - Unauthorized (Invalid/missing token)
- `403` - Forbidden (Insufficient permissions)
- `404` - Not Found
- `429` - Too Many Requests
- `500` - Internal Server Error

## Rate Limiting
- API Rate Limit: 100 requests per minute per IP
- File Upload Rate Limit: 10 uploads per minute per user
- Headers: 
  - X-RateLimit-Limit: 100
  - X-RateLimit-Remaining: 95
  - X-RateLimit-Reset: 1640995200

## System Constraints
- Maximum items per user: 1000
- Maximum categories: 100
- Maximum tags per item: 20
- Image uploads: max 5MB, formats: JPG, PNG
- JWT token expiry: 1 hour
- Refresh token expiry: 30 days