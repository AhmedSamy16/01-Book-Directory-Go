-- +goose Up

CREATE TABLE books (
    id UUID NOT NULL PRIMARY KEY,
    title VARCHAR(120) NOT NULL,
    author VARCHAR(100) NOT NULL
);

-- +goose Down
DROP TABLE books;