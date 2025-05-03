-- name: CreateFeed :one
INSERT into feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds
WHERE user_id = $1;

-- name: GetFeedsWithUsers :many
SELECT
    feeds.id,
    feeds.name,
    feeds.url,
    users.name as user_name
FROM feeds
JOIN users ON feeds.user_id = users.id;