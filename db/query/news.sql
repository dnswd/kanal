-- name: CreateNews :one
INSERT INTO news (
    title,
    author,
    topic,
    status,
    article
) VALUES (
    $1, $2, $3, 'draft', $4
) RETURNING * ;

-- name: AddTagToBlog :one
INSERT INTO news_tag (
    news_id,
    tag_id
) VALUES (
    $1, $2
) RETURNING * ;

-- name: ListNews :many
SELECT n.id, title, t.name, status
FROM news as n
JOIN topic as t ON t.name = n.topic
GROUP BY n.id, t.name;

-- name: ListNewsByTopic :many
SELECT n.id, title, t.name, status
FROM news as n
JOIN topic as t ON t.name = n.topic
WHERE t.name = $1
GROUP BY n.id, t.name;

-- name: ListNewsByTopicAndStatus :many
SELECT n.id, title, t.name, status
FROM news as n
JOIN topic as t ON t.name = n.topic
WHERE t.name = $1 and n.status = $2
GROUP BY n.id, t.name;

-- name: ListNewsByStatus :many
SELECT n.id, title, t.name, status
FROM news as n
JOIN topic as t ON t.name = n.topic
WHERE n.status = $1
GROUP BY n.id, t.name;

-- name: ListNewsByTag :many
SELECT n.id, title, t.name, status
FROM news as n
JOIN topic as t ON t.name = n.topic
WHERE nt.name = (SELECT tag.name from tag where tag.name=$1 LIMIT 1)
GROUP BY n.id, t.name;

-- name: GetNews :one
SELECT n.id, title, t.name, author, published_date, article
FROM news as n
JOIN topic as t ON t.name = n.topic
WHERE n.id = $1 LIMIT 1;

-- name: GetNewsTags :many
SELECT tag.name
FROM news as n, news_tag as nt 
INNER JOIN tag ON nt.tag_id = tag.id
WHERE n.id = $1;

-- name: UpdateNews :exec
UPDATE news SET
    title = $2,
    author = $3,
    topic = $4,
    status = $5,
    published_date = $6,
    article = $7
WHERE id = $1
RETURNING *;

-- name: UpdateNewsStatus :exec
UPDATE news SET
    status = $2
WHERE id = $1
RETURNING *;

-- name: PublishNews :exec
UPDATE news SET
    status = 'published',
    published_date = $2
WHERE id = $1
RETURNING *;

-- name: UnpublishNews :exec
UPDATE news SET
    status = 'deleted'
WHERE id = $1
RETURNING *;

-- name: HardDeleteNews :exec
DELETE FROM news
WHERE id = $1;