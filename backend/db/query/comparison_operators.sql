-- name: GetAllComparisonOperators :many
SELECT id, comparison_operator
FROM comparison_operators;

-- name: GetComparisonOperatorFromID :one
SELECT id, comparison_operator
FROM comparison_operators
WHERE id = ?;
