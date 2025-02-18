-- Insert Categories
INSERT INTO categories (name, description) VALUES
('Meat', 'All types of meat products'),
('Vegetables', 'Fresh and frozen vegetables');

-- Insert Tags
INSERT INTO tags (name) VALUES
('Organic'),
('Quick Prep'),
('Family Size');

-- Insert Items
INSERT INTO items (name, description, packaging, weight_unit, expiration_date) VALUES
('Chicken Breast', 'Organic chicken breast', 'Package', 'kg', CURRENT_DATE + INTERVAL '30 days'),
('Mixed Vegetables', 'Frozen mixed vegetables', 'Bag', 'g', CURRENT_DATE + INTERVAL '90 days'),
('Ground Beef', 'Grass-fed ground beef', 'Package', 'lb', CURRENT_DATE + INTERVAL '45 days');

-- Link Items to Categories
WITH 
    meat_cat AS (SELECT id FROM categories WHERE name = 'Meat'),
    veg_cat AS (SELECT id FROM categories WHERE name = 'Vegetables'),
    chicken AS (SELECT id FROM items WHERE name = 'Chicken Breast'),
    veggies AS (SELECT id FROM items WHERE name = 'Mixed Vegetables'),
    beef AS (SELECT id FROM items WHERE name = 'Ground Beef')
INSERT INTO item_categories (item_id, category_id)
SELECT chicken.id, meat_cat.id FROM chicken, meat_cat
UNION ALL
SELECT veggies.id, veg_cat.id FROM veggies, veg_cat
UNION ALL
SELECT beef.id, meat_cat.id FROM beef, meat_cat;

-- Link Items to Tags
WITH 
    organic_tag AS (SELECT id FROM tags WHERE name = 'Organic'),
    quick_tag AS (SELECT id FROM tags WHERE name = 'Quick Prep'),
    family_tag AS (SELECT id FROM tags WHERE name = 'Family Size'),
    chicken AS (SELECT id FROM items WHERE name = 'Chicken Breast'),
    veggies AS (SELECT id FROM items WHERE name = 'Mixed Vegetables'),
    beef AS (SELECT id FROM items WHERE name = 'Ground Beef')
INSERT INTO item_tags (item_id, tag_id)
SELECT chicken.id, organic_tag.id FROM chicken, organic_tag
UNION ALL
SELECT chicken.id, quick_tag.id FROM chicken, quick_tag
UNION ALL
SELECT veggies.id, quick_tag.id FROM veggies, quick_tag
UNION ALL
SELECT beef.id, family_tag.id FROM beef, family_tag
UNION ALL
SELECT beef.id, organic_tag.id FROM beef, organic_tag;

-- Add some inventory logs
WITH 
    chicken AS (SELECT id FROM items WHERE name = 'Chicken Breast'),
    veggies AS (SELECT id FROM items WHERE name = 'Mixed Vegetables'),
    beef AS (SELECT id FROM items WHERE name = 'Ground Beef')
INSERT INTO inventory_log (item_id, change, weight, weight_unit, notes)
SELECT chicken.id, 5, 2.5, 'kg'::weight_unit, 'Initial stock' FROM chicken
UNION ALL
SELECT veggies.id, 10, 500, 'g'::weight_unit, 'Initial stock' FROM veggies
UNION ALL
SELECT beef.id, 3, 3.0, 'lb'::weight_unit, 'Initial stock' FROM beef; 