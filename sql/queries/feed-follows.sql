-- sqlfluff: disable=L006,L009,L014
-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT 
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feed_follow
INNER JOIN
  feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN 
  users ON users.id = inserted_feed_follow.user_id;
-- sqlfluff: enable=L006,L009,L014

-- name: GetFeedFollowsForUser :many 
SELECT u.name as user_name, f.name as feed_name 
FROM feed_follows as ff
INNER JOIN users u on u.id = ff.user_id
INNER JOIN feeds f on f.id = ff.feed_id
WHERE u.id = $1;


-- name: UnfollowFeed :one
DELETE FROM feed_follows f
WHERE f.user_id = $1 
AND f.feed_id = (
  SELECT feeds.id
  FROM feeds 
  WHERE feeds.url = $2
)
RETURNING *;

