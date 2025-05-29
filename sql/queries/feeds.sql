-- sqlfluff: disable=L006,L009,L014
-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, name, url, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;
-- sqlfluff: enable=L006,L009,L014

-- name: GetFeeds :many
SELECT f.name as feed_name, f.url, u.name as user_name
FROM feeds f
JOIN users u ON f.user_id = u.id; 


-- name: GetFeedByURL :one
SELECT f.id, f.name as feed_name, u.name as user_name, f.url 
FROM feeds f
JOIN users u ON f.user_id = u.id
WHERE f.url = $1;

-- name: MarkFeedFetched :one
UPDATE feeds 
SET updated_at = NOW(),
last_fetched_at = NOW()
WHERE feeds.id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
