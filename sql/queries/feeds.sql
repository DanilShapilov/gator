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

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds 
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;