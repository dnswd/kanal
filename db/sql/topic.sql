-- name: CreateTopic :one
INSERT INTO topic (
    id,
    name
) VALUES (
    $1, $2
) RETURNING *;
