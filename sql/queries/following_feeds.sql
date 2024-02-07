-- name: FollowNewFeed :one
INSERT INTO following_feeds (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ReturnUserFollowingFeeds :many
SELECT * FROM following_feeds WHERE user_id = $1;

-- name: UnfollowFeed :exec
DELETE FROM following_feeds WHERE feed_id = $1 AND user_id = $2;