CREATE TABLE product_categories (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    unit_price NUMERIC NOT NULL,
    unit VARCHAR(50) NOT NULL,
    weight NUMERIC NOT NULL,
    category_id INTEGER REFERENCES product_categories(id),
    status BOOLEAN NOT NULL DEFAULT TRUE
);
