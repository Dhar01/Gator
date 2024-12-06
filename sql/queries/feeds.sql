-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: FetchFeeds :many
SELECT * FROM feeds
JOIN users ON feeds.user_id = users.id
WHERE feeds.user_id = $1
ORDER BY feeds.created_at;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET
    last_fetched_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT id, url
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;