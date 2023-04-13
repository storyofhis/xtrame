CREATE TYPE role AS ENUM (
    'admin',
    'user'
);
-- name: CreateAccounts :one
INSERT INTO accounts (
    fullname, 
    nickname, 
    username,
    email,
    password,
    role,
    age
) VALUES (
    $1, 
    $2, 
    $3, 
    $4, 
    $5, 
    $6, 
    $7
) RETURNING *;

-- name: FindUserByEmail :one
SELECT * FROM accounts 
WHERE email = $1 LIMIT 1;

-- name: FindUserByID :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;