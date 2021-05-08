
CREATE TABLE authors (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    bio  varchar(255)
) ENGINE=InnoDB;

/* name: GetAuthor :one */
SELECT * FROM authors
WHERE id = ? LIMIT 1;

/* name: ListAuthors :many */
SELECT * FROM authors
ORDER BY name LIMIT ?;

/* name: CreateAuthor :exec */
INSERT INTO authors (
          name, bio
) VALUES (
  ?, ?
);

/* name: DeleteAuthor :exec */
DELETE FROM authors
WHERE id = ?;