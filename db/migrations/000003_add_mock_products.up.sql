-- Insert categories
INSERT INTO product_categories (id, category_name) VALUES (1, 'Điện tử') ON CONFLICT DO NOTHING;
INSERT INTO product_categories (id, category_name) VALUES (2, 'Thời trang') ON CONFLICT DO NOTHING;
INSERT INTO product_categories (id, category_name) VALUES (3, 'Văn phòng phẩm') ON CONFLICT DO NOTHING;

-- Insert products
INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES
('Laptop Acer Swift 3', 699.99, 'cái', 2.0, 1, true) ON CONFLICT DO NOTHING;
INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES
('Chuột Bluetooth Silent Ugreen', 15.99, 'cái', 10, 1, true) ON CONFLICT DO NOTHING;
INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES
('iPad Air 5 M1 WiFi 64GB', 499.99, 'cái', 0.2, 2, true) ON CONFLICT DO NOTHING;
INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES
('Xiaomi Mi Band 8', 49.99, 'cái', 0.5, 3, true) ON CONFLICT DO NOTHING;
