-- Disable foreign key checks temporarily for clean deletion
BEGIN;

-- Clean inventory logs first (due to foreign key constraints)
TRUNCATE TABLE inventory_log CASCADE;

-- Clean item relationships
TRUNCATE TABLE item_categories CASCADE;
TRUNCATE TABLE item_tags CASCADE;

-- Clean main tables
TRUNCATE TABLE items CASCADE;
TRUNCATE TABLE categories CASCADE;
TRUNCATE TABLE tags CASCADE;

-- Re-enable foreign key checks
COMMIT; 