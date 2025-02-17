-- name: WriteUUID :one
INSERT INTO payments(id, price, is_enabled)
VALUES ($1, $2, $3)
RETURNING *;