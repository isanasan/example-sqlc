-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;


-- name: CreateUser :execresult
INSERT INTO users (
  name, email, age
) VALUES (
  $1, $2, $3
);

-- name: UpdateUserAges :exec
UPDATE users SET age = $2
WHERE id = $1;
