ALTER TABLE products ADD COLUMN product_search tsvector;

-- Create a GIN index on the search_vector column
CREATE INDEX products_search_vector_idx ON products USING GIN (product_search);

-- Update the search_vector column with existing data
UPDATE products SET product_search = to_tsvector('simple', product_name);

-- Create a trigger to update the search_vector column on insert or update
CREATE TRIGGER products_search_vector_update
BEFORE INSERT OR UPDATE ON products
FOR EACH ROW EXECUTE FUNCTION tsvector_update_trigger('product_search', 'pg_catalog.simple', 'product_name');
