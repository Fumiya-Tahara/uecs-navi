// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: uecs_devices.sql

package mysqlc

import (
	"context"
)

const createUecsDevice = `-- name: CreateUecsDevice :execlastid
INSERT INTO uecs_devices (ccmtype, room, region, ` + "`" + `order` + "`" + `, ` + "`" + `priority` + "`" + `)
VALUES (?, ?, ?, ?, ?)
`

type CreateUecsDeviceParams struct {
	Ccmtype  string
	Room     int32
	Region   int32
	Order    int32
	Priority int32
}

func (q *Queries) CreateUecsDevice(ctx context.Context, arg CreateUecsDeviceParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createUecsDevice,
		arg.Ccmtype,
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

const getAllUecsDevice = `-- name: GetAllUecsDevice :many
SELECT id, ccmtype, room, region, ` + "`" + `order` + "`" + `, ` + "`" + `priority` + "`" + `
FROM uecs_devices
`

type GetAllUecsDeviceRow struct {
	ID       int32
	Ccmtype  string
	Room     int32
	Region   int32
	Order    int32
	Priority int32
}

func (q *Queries) GetAllUecsDevice(ctx context.Context) ([]GetAllUecsDeviceRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllUecsDevice)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUecsDeviceRow
	for rows.Next() {
		var i GetAllUecsDeviceRow
		if err := rows.Scan(
			&i.ID,
			&i.Ccmtype,
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

const getUecsDeviceFromID = `-- name: GetUecsDeviceFromID :one
SELECT id, ccmtype, room, region, ` + "`" + `order` + "`" + `, ` + "`" + `priority` + "`" + `
FROM uecs_devices
WHERE id = ?
`

type GetUecsDeviceFromIDRow struct {
	ID       int32
	Ccmtype  string
	Room     int32
	Region   int32
	Order    int32
	Priority int32
}

func (q *Queries) GetUecsDeviceFromID(ctx context.Context, id int32) (GetUecsDeviceFromIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUecsDeviceFromID, id)
	var i GetUecsDeviceFromIDRow
	err := row.Scan(
		&i.ID,
		&i.Ccmtype,
		&i.Room,
		&i.Region,
		&i.Order,
		&i.Priority,
	)
	return i, err
}