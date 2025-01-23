// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: workflow_operations.sql

package mysqlc

import (
	"context"
	"database/sql"
)

const createWorkflowOperation = `-- name: CreateWorkflowOperation :execlastid
INSERT INTO workflow_operations (workflow_id, relay_1, relay_2, relay_3, relay_4, relay_5, relay_6, relay_7, relay_8) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateWorkflowOperationParams struct {
	WorkflowID int32
	Relay1     sql.NullBool
	Relay2     sql.NullBool
	Relay3     sql.NullBool
	Relay4     sql.NullBool
	Relay5     sql.NullBool
	Relay6     sql.NullBool
	Relay7     sql.NullBool
	Relay8     sql.NullBool
}

func (q *Queries) CreateWorkflowOperation(ctx context.Context, arg CreateWorkflowOperationParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createWorkflowOperation,
		arg.WorkflowID,
		arg.Relay1,
		arg.Relay2,
		arg.Relay3,
		arg.Relay4,
		arg.Relay5,
		arg.Relay6,
		arg.Relay7,
		arg.Relay8,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getWorkflowOperationsFromWorkflow = `-- name: GetWorkflowOperationsFromWorkflow :one
SELECT id, workflow_id, relay_1, relay_2, relay_3, relay_4, relay_5, relay_6, relay_7, relay_8
FROM workflow_operations
WHERE workflow_id = ?
`

type GetWorkflowOperationsFromWorkflowRow struct {
	ID         int32
	WorkflowID int32
	Relay1     sql.NullBool
	Relay2     sql.NullBool
	Relay3     sql.NullBool
	Relay4     sql.NullBool
	Relay5     sql.NullBool
	Relay6     sql.NullBool
	Relay7     sql.NullBool
	Relay8     sql.NullBool
}

func (q *Queries) GetWorkflowOperationsFromWorkflow(ctx context.Context, workflowID int32) (GetWorkflowOperationsFromWorkflowRow, error) {
	row := q.db.QueryRowContext(ctx, getWorkflowOperationsFromWorkflow, workflowID)
	var i GetWorkflowOperationsFromWorkflowRow
	err := row.Scan(
		&i.ID,
		&i.WorkflowID,
		&i.Relay1,
		&i.Relay2,
		&i.Relay3,
		&i.Relay4,
		&i.Relay5,
		&i.Relay6,
		&i.Relay7,
		&i.Relay8,
	)
	return i, err
}
