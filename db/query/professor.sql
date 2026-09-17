-- name: CreateProfessor :one
INSERT INTO professors (
  first_name, last_name, email, department_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetProfessor :one
SELECT * FROM professors
WHERE id = $1 LIMIT 1;

-- name: ListProfessors :many
SELECT * FROM professors
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: DeleteProfessor :exec
DELETE FROM professors
WHERE id = $1;