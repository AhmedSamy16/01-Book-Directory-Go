-- name: GetAllBooks :many
SELECT * FROM books;

-- name: CreateBook :one
INSERT INTO books(id, title, author)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetBookById :one
SELECT * FROM books WHERE id = $1;

-- name: UpdateBookById :one
UPDATE books SET title = $1, author = $2 WHERE id = $3
RETURNING *;

-- name: DeleteBookById :one
DELETE FROM books WHERE id = $1
RETURNING *;