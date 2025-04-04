// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: conditions.sql

package mysqlc

import (
	"context"
)

const createCondition = `-- name: CreateCondition :execlastid
INSERT INTO conditions (climate_data_id, time_schedule_row_id, comparison_operator_id, set_point)
VALUES (?, ?, ?, ?)
`

type CreateConditionParams struct {
	ClimateDataID        int32
	TimeScheduleRowID    int32
	ComparisonOperatorID int32
	SetPoint             float64
}

func (q *Queries) CreateCondition(ctx context.Context, arg CreateConditionParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createCondition,
		arg.ClimateDataID,
		arg.TimeScheduleRowID,
		arg.ComparisonOperatorID,
		arg.SetPoint,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getConditionFromTimeScheduleRow = `-- name: GetConditionFromTimeScheduleRow :one
SELECT id, climate_data_id, time_schedule_row_id, comparison_operator_id, set_point
FROM conditions
WHERE time_schedule_row_id = ?
`

type GetConditionFromTimeScheduleRowRow struct {
	ID                   int32
	ClimateDataID        int32
	TimeScheduleRowID    int32
	ComparisonOperatorID int32
	SetPoint             float64
}

func (q *Queries) GetConditionFromTimeScheduleRow(ctx context.Context, timeScheduleRowID int32) (GetConditionFromTimeScheduleRowRow, error) {
	row := q.db.QueryRowContext(ctx, getConditionFromTimeScheduleRow, timeScheduleRowID)
	var i GetConditionFromTimeScheduleRowRow
	err := row.Scan(
		&i.ID,
		&i.ClimateDataID,
		&i.TimeScheduleRowID,
		&i.ComparisonOperatorID,
		&i.SetPoint,
	)
	return i, err
}

const updateCondition = `-- name: UpdateCondition :exec
UPDATE conditions
SET climate_data_id = ?, time_schedule_row_id = ?, comparison_operator_id = ?, set_point = ?
WHERE id = ?
`

type UpdateConditionParams struct {
	ClimateDataID        int32
	TimeScheduleRowID    int32
	ComparisonOperatorID int32
	SetPoint             float64
	ID                   int32
}

func (q *Queries) UpdateCondition(ctx context.Context, arg UpdateConditionParams) error {
	_, err := q.db.ExecContext(ctx, updateCondition,
		arg.ClimateDataID,
		arg.TimeScheduleRowID,
		arg.ComparisonOperatorID,
		arg.SetPoint,
		arg.ID,
	)
	return err
}
