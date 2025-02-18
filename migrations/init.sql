-- Create database (run this separately if database doesn't exist)
CREATE DATABASE freezer_db;

-- Connect to the database
\c freezer_db;

-- Create ENUM types
CREATE TYPE user_role AS ENUM ('admin', 'user');
CREATE TYPE weight_unit AS ENUM ('g', 'kg', 'oz', 'lb');

-- Create tables
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    barcode VARCHAR(50),
    image_url TEXT,
    packaging VARCHAR(50),
    weight_unit weight_unit,
    expiration_date DATE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE inventory_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id UUID NOT NULL REFERENCES items(id),
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    change INTEGER NOT NULL,
    weight FLOAT,
    weight_unit weight_unit,
    notes TEXT
);

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE item_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id UUID NOT NULL REFERENCES items(id),
    category_id UUID NOT NULL REFERENCES categories(id),
    UNIQUE(item_id, category_id)
);

CREATE TABLE item_tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id UUID NOT NULL REFERENCES items(id),
    tag_id UUID NOT NULL REFERENCES tags(id),
    UNIQUE(item_id, tag_id)
);

-- Create indexes
CREATE INDEX idx_items_name ON items(name);
CREATE INDEX idx_inventory_log_item_id ON inventory_log(item_id);
CREATE INDEX idx_item_categories_item_id ON item_categories(item_id);
CREATE INDEX idx_item_tags_item_id ON item_tags(item_id);

-- Insert admin user (password: admin123)
INSERT INTO users (email, password_hash, role) VALUES 
('admin@example.com', '$2a$10$ZkX5nxV1cJ/CD9UW9LC37OmpxS99p3eQBJ.1kV.bm5/QwZ.fx.OXe', 'admin'); 