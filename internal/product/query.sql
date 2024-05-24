-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProductByTitle :one
SELECT * FROM products
WHERE title = $1 LIMIT 1;

-- name: GetProducts :many
SELECT * FROM products
WHERE id > $1 LIMIT $2;

-- name: GetProductsBySize :many
SELECT * FROM products
WHERE size = $1 OR size_eu = $2 OR size_us = $3;

-- name: GetProductsFromDate :many
SELECT * FROM products
WHERE created_at >= $1;

-- name: CreateProduct :one
INSERT INTO products (
  title, description, img, size, size_eu, size_us, price, sex, kind, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;