-- name: GetDepartment :one
SELECT * FROM departments
WHERE id = $1 LIMIT 1;

-- name: ListDepartments :many
SELECT * FROM departments
ORDER BY id;

-- name: CreateDepartment :one
INSERT INTO departments (
  name
) VALUES (
  $1
)
RETURNING *;