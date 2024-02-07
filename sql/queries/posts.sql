-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, description, published_at, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
WITH ranked_posts AS (
    SELECT posts.*,
           ROW_NUMBER() OVER (PARTITION BY posts.feed_id ORDER BY posts.published_at DESC) AS rn, feeds.name AS feed
    FROM posts
    JOIN following_feeds ON posts.feed_id = following_feeds.feed_id
	JOIN feeds ON posts.feed_id = feeds.id
    WHERE following_feeds.user_id = $1
)
SELECT * FROM ranked_posts
WHERE rn <=  5
ORDER BY published_at DESC;