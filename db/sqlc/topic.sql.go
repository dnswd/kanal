// Code generated by sqlc. DO NOT EDIT.
// source: topic.sql

package db

import (
	"context"
)

const createTopic = `-- name: CreateTopic :exec
INSERT INTO topic (
    name
) VALUES (
    $1
) RETURNING id, name
`

func (q *Queries) CreateTopic(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, createTopic, name)
	return err
}

const listTopic = `-- name: ListTopic :many
SELECT id, name
FROM topic
`

func (q *Queries) ListTopic(ctx context.Context) ([]Topic, error) {
	rows, err := q.db.QueryContext(ctx, listTopic)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Topic
	for rows.Next() {
		var i Topic
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const renameTopic = `-- name: RenameTopic :exec
UPDATE topic SET
    name = $2
WHERE name = $1
RETURNING id, name
`

type RenameTopicParams struct {
	Name   string `json:"name"`
	Name_2 string `json:"name_2"`
}

func (q *Queries) RenameTopic(ctx context.Context, arg RenameTopicParams) error {
	_, err := q.db.ExecContext(ctx, renameTopic, arg.Name, arg.Name_2)
	return err
}
