-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- use "docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate" instead of "sqlc generate" you might need to run "docker pull kjconroy/sqlc"
