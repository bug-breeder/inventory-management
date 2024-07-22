ALTER TABLE products ADD COLUMN image_url VARCHAR(255);
ALTER TABLE products ADD COLUMN description TEXT;

-- Update existing products with description and image_url
UPDATE products SET
  description = 'A high-performance laptop with Intel Core i5 processor and 8GB RAM.',
  image_url = 'public/img1.jpg'
WHERE product_name = 'Laptop Acer Swift 3';

UPDATE products SET
  description = 'A silent Bluetooth mouse, perfect for noise-free operation.',
  image_url = 'public/img2.jpg'
WHERE product_name = 'Chuột Bluetooth Silent Ugreen';

UPDATE products SET
  description = 'The latest iPad Air with M1 chip and 64GB storage, WiFi model.',
  image_url = 'public/img3.jpg'
WHERE product_name = 'iPad Air 5 M1 WiFi 64GB';

UPDATE products SET
  description = 'A fitness tracker with advanced health monitoring features.',
  image_url = 'public/img4.jpg'
WHERE product_name = 'Xiaomi Mi Band 8';

