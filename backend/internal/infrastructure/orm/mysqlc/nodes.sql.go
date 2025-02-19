// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: nodes.sql

package mysqlc

import (
	"context"
	"encoding/json"
)

const createNode = `-- name: CreateNode :execlastid
INSERT INTO nodes (workflow_id, workflow_node_id, type, data, position_x, position_y) 
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateNodeParams struct {
	WorkflowID     int32
	WorkflowNodeID string
	Type           string
	Data           json.RawMessage
	PositionX      float64
	PositionY      float64
}

func (q *Queries) CreateNode(ctx context.Context, arg CreateNodeParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createNode,
		arg.WorkflowID,
		arg.WorkflowNodeID,
		arg.Type,
		arg.Data,
		arg.PositionX,
		arg.PositionY,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getNodesFromWorkflow = `-- name: GetNodesFromWorkflow :many
SELECT 
    id, workflow_id, workflow_node_id, type, data, position_x, position_y
FROM nodes
WHERE workflow_id = ?
`

type GetNodesFromWorkflowRow struct {
	ID             int32
	WorkflowID     int32
	WorkflowNodeID string
	Type           string
	Data           json.RawMessage
	PositionX      float64
	PositionY      float64
}

func (q *Queries) GetNodesFromWorkflow(ctx context.Context, workflowID int32) ([]GetNodesFromWorkflowRow, error) {
	rows, err := q.db.QueryContext(ctx, getNodesFromWorkflow, workflowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNodesFromWorkflowRow
	for rows.Next() {
		var i GetNodesFromWorkflowRow
		if err := rows.Scan(
			&i.ID,
			&i.WorkflowID,
			&i.WorkflowNodeID,
			&i.Type,
			&i.Data,
			&i.PositionX,
			&i.PositionY,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
