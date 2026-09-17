-- name: CreateStudent :one
INSERT INTO students (
  first_name, last_name, email, department_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: ListStudents :many
SELECT * FROM students
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetStudent :one
SELECT * FROM students
WHERE id = $1 LIMIT 1;