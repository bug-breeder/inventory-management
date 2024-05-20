CREATE EXTENSION unaccent;

ALTER TABLE products ADD COLUMN product_search tsvector;

-- Create a GIN index on the search_vector column
CREATE INDEX products_search_vector_idx ON products USING GIN (product_search);

-- Create a trigger to update the search_vector column on insert or update using unaccent
CREATE OR REPLACE FUNCTION update_search_vector() RETURNS trigger AS $$
BEGIN
  NEW.product_search := to_tsvector('simple', unaccent(NEW.product_name));
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER products_search_vector_update
BEFORE INSERT OR UPDATE ON products
FOR EACH ROW EXECUTE FUNCTION update_search_vector();
