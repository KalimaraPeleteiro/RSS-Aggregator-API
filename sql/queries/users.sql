-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name, api_key)
VALUES($1, $2, $3, $4,
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUseByAPIKey :one
SELECT * FROM users where api_key = $1;

-- name: LoginAuthentication :one
SELECT * FROM users where name = $1 and password = $2
LIMIT 1;