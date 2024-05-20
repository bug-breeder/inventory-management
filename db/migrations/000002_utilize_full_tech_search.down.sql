-- Drop the trigger
DROP TRIGGER IF EXISTS products_search_vector_update ON products;

-- Drop the function
DROP FUNCTION IF EXISTS update_search_vector();

-- Drop the index
DROP INDEX IF EXISTS products_search_vector_idx;

-- Drop the column
ALTER TABLE products DROP COLUMN IF EXISTS product_search;

-- Drop the extension
DROP EXTENSION IF EXISTS unaccent;
