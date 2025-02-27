-- name: CreateTask :one
INSERT INTO tasks (id, name, is_completed, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTaskByID :one
SELECT * FROM tasks WHERE id = $1;

-- name: GetAllTasks :many
SELECT * FROM tasks ORDER BY created_at DESC;

-- name: UpdateTaskStatus :exec
UPDATE tasks SET is_completed = $2, updated_at = $3 WHERE id = $1;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;
