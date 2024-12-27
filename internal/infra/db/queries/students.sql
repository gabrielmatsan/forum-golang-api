-- name: CreateStudent :exec
INSERT INTO students (
  id,
  name,
  email,
  password
) VALUES (
  $1,
  $2,
  $3,
  $4
);

-- name: FindByEmail :one
SELECT * FROM students WHERE email = $1;

-- name: FindById :one
SELECT * FROM students WHERE id = $1;

-- name: UpdateStudent :exec
UPDATE students SET name = $2, email = $3, password = $4 WHERE id = $1;