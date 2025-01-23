-- name: CreateWorkflowOperation :execlastid
INSERT INTO workflow_operations (workflow_id, relay_1, relay_2, relay_3, relay_4, relay_5, relay_6, relay_7, relay_8) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetWorkflowOperationsFromWorkflow :one
SELECT id, workflow_id, relay_1, relay_2, relay_3, relay_4, relay_5, relay_6, relay_7, relay_8
FROM workflow_operations
WHERE workflow_id = ?;
