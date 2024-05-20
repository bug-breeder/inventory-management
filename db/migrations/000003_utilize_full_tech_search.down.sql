-- Drop the trigger that updates the search_vector column
DROP TRIGGER IF EXISTS products_search_vector_update ON products;

-- Drop the GIN index on the search_vector column
DROP INDEX IF EXISTS products_search_vector_idx;

-- Drop the search_vector column
ALTER TABLE products DROP COLUMN IF EXISTS product_search;
