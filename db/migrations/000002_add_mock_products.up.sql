-- Insert categories
INSERT INTO product_categories (id, category_name) VALUES (1, 'Điện tử') ON CONFLICT DO NOTHING;
INSERT INTO product_categories (id, category_name) VALUES (2, 'Thời trang') ON CONFLICT DO NOTHING;
INSERT INTO product_categories (id, category_name) VALUES (3, 'Văn phòng phẩm') ON CONFLICT DO NOTHING;

-- Insert products
INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES
('Laptop Acer Swift 3', 699.99, 'cái', 2.0, 1, true) ON CONFLICT DO NOTHING;
INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES
('Pc Phong Vũ', 1299.99, 'cái', 10, 1, true) ON CONFLICT DO NOTHING;
INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES
('Aó khoác nam', 9.99, 'cái', 0.2, 2, true) ON CONFLICT DO NOTHING;
INSERT INTO products (product_name, unit_price, unit, weight, category_id, status) VALUES
('Sổ tay Teko', 19.99, 'quyển', 0.5, 3, true) ON CONFLICT DO NOTHING;
