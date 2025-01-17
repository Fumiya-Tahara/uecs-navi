// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: devices.sql

package mysqlc

import (
	"context"
	"database/sql"
)

const createDevice = `-- name: CreateDevice :execlastid
INSERT INTO devices (climate_data_id, m304_id, name, rly) 
VALUES (?, ?, ?, ?)
`

type CreateDeviceParams struct {
	ClimateDataID int32
	M304ID        int32
	Name          string
	Rly           sql.NullInt32
}

func (q *Queries) CreateDevice(ctx context.Context, arg CreateDeviceParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createDevice,
		arg.ClimateDataID,
		arg.M304ID,
		arg.Name,
		arg.Rly,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getDevicesFromM304 = `-- name: GetDevicesFromM304 :many
SELECT 
    id, climate_data_id, m304_id, name, rly
FROM devices
WHERE m304_id = ?
`

type GetDevicesFromM304Row struct {
	ID            int32
	ClimateDataID int32
	M304ID        int32
	Name          string
	Rly           sql.NullInt32
}

func (q *Queries) GetDevicesFromM304(ctx context.Context, m304ID int32) ([]GetDevicesFromM304Row, error) {
	rows, err := q.db.QueryContext(ctx, getDevicesFromM304, m304ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDevicesFromM304Row
	for rows.Next() {
		var i GetDevicesFromM304Row
		if err := rows.Scan(
			&i.ID,
			&i.ClimateDataID,
			&i.M304ID,
			&i.Name,
			&i.Rly,
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
