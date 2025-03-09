-- name: CreateEdge :execlastid
INSERT INTO edges (workflow_id, source_node_id, target_node_id) 
VALUES (?, ?, ?);

-- name: GetEdgesFromWorkflow :many
SELECT 
    id, workflow_id, source_node_id, target_node_id
FROM edges
WHERE workflow_id = ?;

-- name: UpdateEdge :exec
UPDATE edges
SET workflow_id = ?, source_node_id = ?, target_node_id = ?
WHERE id = ?;
