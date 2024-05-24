CREATE DATABASE le_donne;

\c le_donne;

CREATE TABLE users (
  id         BIGSERIAL PRIMARY KEY,
  email      VARCHAR(255) UNIQUE NOT NULL,
  username   VARCHAR(32) UNIQUE NOT NULL,
  password   VARCHAR(72) NOT NULL,
  created_at TIMESTAMP,
  admin      BOOLEAN
);

CREATE TABLE product_kind (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL
);

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