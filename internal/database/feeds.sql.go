// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feeds.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id, created_at, updated_at, last_fetched_at, name, url, user_id
`

type CreateFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastFetchedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}

const deleteAllFeeds = `-- name: DeleteAllFeeds :exec
DELETE FROM feeds
`

func (q *Queries) DeleteAllFeeds(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllFeeds)
	return err
}

const fetchFeeds = `-- name: FetchFeeds :many
SELECT feeds.id, feeds.created_at, feeds.updated_at, last_fetched_at, feeds.name, url, user_id, users.id, users.created_at, users.updated_at, users.name FROM feeds
JOIN users ON feeds.user_id = users.id
WHERE feeds.user_id = $1
ORDER BY feeds.created_at
`

type FetchFeedsRow struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	LastFetchedAt sql.NullTime
	Name          string
	Url           string
	UserID        uuid.UUID
	ID_2          uuid.UUID
	CreatedAt_2   time.Time
	UpdatedAt_2   time.Time
	Name_2        string
}

func (q *Queries) FetchFeeds(ctx context.Context, userID uuid.UUID) ([]FetchFeedsRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchFeeds, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchFeedsRow
	for rows.Next() {
		var i FetchFeedsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LastFetchedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.Name_2,
		); err != nil {
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

const getAllFeeds = `-- name: GetAllFeeds :many
SELECT feeds.id, feeds.created_at, feeds.updated_at, last_fetched_at, feeds.name, url, user_id, users.id, users.created_at, users.updated_at, users.name FROM feeds
JOIN users ON feeds.user_id = users.id
`

type GetAllFeedsRow struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	LastFetchedAt sql.NullTime
	Name          string
	Url           string
	UserID        uuid.UUID
	ID_2          uuid.UUID
	CreatedAt_2   time.Time
	UpdatedAt_2   time.Time
	Name_2        string
}

func (q *Queries) GetAllFeeds(ctx context.Context) ([]GetAllFeedsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllFeedsRow
	for rows.Next() {
		var i GetAllFeedsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LastFetchedAt,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.Name_2,
		); err != nil {
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

const getFeedByURL = `-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, last_fetched_at, name, url, user_id FROM feeds
WHERE url = $1
`

func (q *Queries) GetFeedByURL(ctx context.Context, url string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByURL, url)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastFetchedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}
