ALTER TABLE products
DROP COLUMN IF EXISTS image_url,
DROP COLUMN IF EXISTS description;

-- Add image_url and description columns to the products table
ALTER TABLE products
ADD COLUMN image_url VARCHAR(255),
ADD COLUMN description TEXT;

