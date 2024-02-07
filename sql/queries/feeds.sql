-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.*,
       CASE WHEN following_feeds.user_id IS NOT NULL THEN TRUE ELSE FALSE END AS is_following
FROM feeds
LEFT JOIN following_feeds ON feeds.id = following_feeds.feed_id AND following_feeds.user_id = $1;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_time_fetched ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds 
SET last_time_fetched = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;