-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListFeeds :many
SELECT feeds.name, url, users.name as username
	FROM feeds as feeds
	inner join users as users
		on feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECT * FROM feeds 
WHERE url = $1;