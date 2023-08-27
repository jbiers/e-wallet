-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY username;

-- name: CreateAccount :one
INSERT INTO accounts (
    username, type, document, email, password
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: UpdateAccount :one
UPDATE accounts
  set balance = $2
WHERE id = $1
RETURNING *;
