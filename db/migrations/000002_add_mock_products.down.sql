-- Delete products
DELETE FROM products WHERE product_name IN ('Laptop Acer Swift 3', 'Pc Phong Vũ 1', 'Aó khoác nam', 'Sổ tay Teko');

-- Delete categories
DELETE FROM product_categories WHERE id IN (1, 2, 3);
