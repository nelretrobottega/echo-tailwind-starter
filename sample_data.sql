INSERT INTO product_kind (title, description) VALUES ('scarpe', 'hello');

INSERT INTO products (
  title, description, img, size, size_eu, size_us, price, sex, kind, created_at
) VALUES (
  'Scarpa 1', 
  'Ok ok', 
  'https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.jpg', 
  27, 
  42, 
  9, 
  120, 
  0, 
  1, 
  NOW()
);