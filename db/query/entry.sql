-- name: CreateEntry :one
INSERT INTO entries (
    account_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1
LIMIT 1;

-- name: ListAllEntriesForAccount :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListAllEntries :many
SELECT * FROM entries
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: UpdateEntry :exec
UPDATE entries SET amount = $2
WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1;