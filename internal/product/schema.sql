CREATE TABLE products (
  id BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMP,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  img VARCHAR(255) NOT NULL,
  size INTEGER NOT NULL,
  size_us INTEGER,
  size_eu INTEGER,
  price DECIMAL NOT NULL,
  sex INTEGER,
  kind SERIAL REFERENCES product_kind(id)
);