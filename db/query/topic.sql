-- name: CreateTopic :exec
INSERT INTO topic (
    name
) VALUES (
    $1
) RETURNING *;

-- name: ListTopic :many
SELECT id, name
FROM topic;

-- name: RenameTopic :exec
UPDATE topic SET
    name = $2
WHERE name = $1
RETURNING *;