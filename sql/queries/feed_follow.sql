-- name: CreateFeedFollow :one
WITH inserted_feed_follow as (INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *)
select *, feeds.name as feed_name, users.name as user_name from inserted_feed_follow
inner join feeds on inserted_feed_follow.feed_id = feeds.id
inner join users on inserted_feed_follow.user_id = users.id
;

-- name: GetFeedsByUserId :many
SELECT feed_follows.*, feeds.name as feed_name, users.name as user_name FROM feed_follows
inner join feeds on feed_follows.feed_id = feeds.id
inner join users on feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;