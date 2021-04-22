-- name: AddTag :one
INSERT INTO tag (
    name
) VALUES (
    $1
) REUTNING *;

-- name: ListTag :many
SELECT id, name
FROM tag
ORDER BY id ASC;

-- name: ListName :many
SELECT id, name
FROM tag
ORDER BY name ASC;

-- name: GetTagById :one
SELECT id, name
FROM tag
WHERE id = $1;

-- name: RenameTag :exec
UPDATE tag SET
    name = $2
WHERE id = $1;

-- name: DeleteTag :exec
DELETE FROM tag
WHERE id = $1;