-- name: CreateWorkflow :execlastid
INSERT INTO workflows (m304_id, name) 
VALUES (?, ?);

-- name: GetWorkflowsFromM304 :many
SELECT id, m304_id, name
FROM workflows
WHERE m304_id = ?;

-- name: UpdateWorkflow :exec
UPDATE workflows
SET m304_id = ?, name = ?
WHERE id = ?;

-- name: DeleteWorkflow :exec
DELETE FROM workflows
WHERE id = ?;
