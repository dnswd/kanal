-- name: CreateTopic :exec
INSERT INTO topic (
    name
) VALUES (
    $1
) RETURNING *;

-- name: ListTopic :many
SELECT id, name
FROM topic;