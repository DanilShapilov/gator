-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id, created_at, updated_at)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE name = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;