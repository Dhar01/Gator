-- name: CreatePost :one
INSERT INTO posts(id, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetPostsByUser :many
SELECT posts.title, posts.url, posts.published_at
FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.updated_at DESC;


-- name: DeleteAllPosts :exec
DELETE FROM posts;