-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, created_at, updated_at, name, url)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT id, user_id, created_at, updated_at, name, url, last_fetched_at
FROM feeds;

-- name: GetFeedByURL :one
SELECT id, user_id, created_at, updated_at, name, url, last_fetched_at
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :one 
UPDATE feeds
SET
    updated_at = NOW(),
    last_fetched_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;