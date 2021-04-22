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

-- name: ListNews :many
SELECT n.id, title, t.name, status, array_agg(nt.name)
FROM news as n
JOIN topic as t ON t.name = n.topic
LEFT JOIN news_tag as nt ON nt.news_id = n.id
GROUP BY n.id;

-- name: ListNewsByTopic :many
SELECT n.id, title, t.name, status, array_agg(nt.name)
FROM news as n
JOIN topic as t ON t.name = n.topic
LEFT JOIN news_tag as nt ON nt.news_id = n.id
WHERE t.id = $1
GROUP BY n.id;

-- name: ListNewsByStatus :many
SELECT n.id, title, t.name, status, array_agg(nt.name)
FROM news as n
JOIN topic as t ON t.name = n.topic
LEFT JOIN news_tag as nt ON nt.news_id = n.id
WHERE status = $1
GROUP BY n.id;

-- name: ListNewsByTag :many
SELECT n.id, title, t.name, status
FROM news as n
JOIN topic as t ON t.name = n.topic
LEFT JOIN news_tag as nt ON nt.news_id = n.id
WHERE nt.name = (SELECT tag.name from tag where tag.name=$1 LIMIT 1)
GROUP BY n.id;

-- name: GetNews :one
SELECT n.id, title, t.name, array_agg(nt.name), author, published_date, article
FROM news as n
JOIN topic as t ON t.name = n.topic
LEFT JOIN news_tag as nt ON nt.news_id = n.id
WHERE n.id = $1 LIMIT 1;

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

-- name: DeleteNews :exec
UPDATE news SET
    status = 'deleted'
WHERE id = $1
RETURNING *;
