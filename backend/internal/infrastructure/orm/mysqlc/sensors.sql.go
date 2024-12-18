// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sensors.sql

package mysqlc

import (
	"context"
)

const createSensor = `-- name: CreateSensor :execlastid
INSERT INTO sensors (ccm_type, room, region, ` + "`" + `order` + "`" + `, ` + "`" + `priority` + "`" + `)
VALUES (?, ?, ?, ?, ?)
`

type CreateSensorParams struct {
	CcmType  string
	Room     int32
	Region   int32
	Order    int32
	Priority int32
}

func (q *Queries) CreateSensor(ctx context.Context, arg CreateSensorParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createSensor,
		arg.CcmType,
		arg.Room,
		arg.Region,
		arg.Order,
		arg.Priority,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getAllSensor = `-- name: GetAllSensor :many
SELECT id, ccm_type, room, region, ` + "`" + `order` + "`" + `, ` + "`" + `priority` + "`" + `
FROM sensors
`

type GetAllSensorRow struct {
	ID       int32
	CcmType  string
	Room     int32
	Region   int32
	Order    int32
	Priority int32
}

func (q *Queries) GetAllSensor(ctx context.Context) ([]GetAllSensorRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllSensor)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllSensorRow
	for rows.Next() {
		var i GetAllSensorRow
		if err := rows.Scan(
			&i.ID,
			&i.CcmType,
			&i.Room,
			&i.Region,
			&i.Order,
			&i.Priority,
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

const getSensorFromID = `-- name: GetSensorFromID :one
SELECT id, ccm_type, room, region, ` + "`" + `order` + "`" + `, ` + "`" + `priority` + "`" + `
FROM sensors
WHERE id = ?
`

type GetSensorFromIDRow struct {
	ID       int32
	CcmType  string
	Room     int32
	Region   int32
	Order    int32
	Priority int32
}

func (q *Queries) GetSensorFromID(ctx context.Context, id int32) (GetSensorFromIDRow, error) {
	row := q.db.QueryRowContext(ctx, getSensorFromID, id)
	var i GetSensorFromIDRow
	err := row.Scan(
		&i.ID,
		&i.CcmType,
		&i.Room,
		&i.Region,
		&i.Order,
		&i.Priority,
	)
	return i, err
}
