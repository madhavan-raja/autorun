-- name: GetProcess :one
SELECT * FROM process
WHERE id = ? LIMIT 1;

-- name: ListProcesses :many
SELECT * FROM process
ORDER BY id;

-- name: AddProcess :one
INSERT INTO process (
  name, description, command, interval
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: DeleteProcess :exec
DELETE FROM process
WHERE id = ?;