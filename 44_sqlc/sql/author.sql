CREATE TABLE authors (
    id int NOT NULL,
    name varchar(255) NOT NULL,
    bio  varchar(255) NOT NULL
);

/* name: GetAuthor :one */
SELECT * FROM authors
WHERE id = ? LIMIT 1;

/* name: ListAuthors :many */
SELECT * FROM authors
ORDER BY name LIMIT ? OFFSET ?;

/* name: CreateAuthor :exec */
INSERT INTO authors (
          name, bio
) VALUES (
  ?, ?
);

/* name: LastInsertId :one */
SELECT LAST_INSERT_ID();

/* name: RowsAffected :one */
SELECT ROW_COUNT();

/* name: DeleteAuthor :exec */
DELETE FROM authors
WHERE id = ?;

/* name: UpdateAuthor :exec */
UPDATE authors set name = ?, bio = ? WHERE id = ?;

/* name: ListAuthorID :many */
SELECT id FROM authors
ORDER BY id LIMIT ?;

