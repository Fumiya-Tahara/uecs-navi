// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: devices.sql

package mysqlc

import (
	"context"
	"database/sql"
	"time"
)

const createDevice = `-- name: CreateDevice :execlastid
INSERT INTO devices (house_id, climate_data_id, device_name, set_point, duration) 
VALUES (?, ?, ?, ?, ?)
`

type CreateDeviceParams struct {
	HouseID       int32
	ClimateDataID int32
	DeviceName    sql.NullString
	SetPoint      sql.NullFloat64
	Duration      sql.NullInt32
}

func (q *Queries) CreateDevice(ctx context.Context, arg CreateDeviceParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createDevice,
		arg.HouseID,
		arg.ClimateDataID,
		arg.DeviceName,
		arg.SetPoint,
		arg.Duration,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getDevicesFromHouse = `-- name: GetDevicesFromHouse :many
SELECT 
    id, house_id, climate_data_id, device_name, set_point, duration, created_at, updated_at
FROM devices
WHERE house_id = ?
`

func (q *Queries) GetDevicesFromHouse(ctx context.Context, houseID int32) ([]Device, error) {
	rows, err := q.db.QueryContext(ctx, getDevicesFromHouse, houseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Device
	for rows.Next() {
		var i Device
		if err := rows.Scan(
			&i.ID,
			&i.HouseID,
			&i.ClimateDataID,
			&i.DeviceName,
			&i.SetPoint,
			&i.Duration,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getJoinedDevicesFromHouse = `-- name: GetJoinedDevicesFromHouse :many
SELECT 
    d.id, d.house_id, d.device_name, d.set_point, d.duration, d.created_at, d.updated_at,
    c.name AS climate_data_name, c.unit
FROM devices d
JOIN climate_datas c ON d.climate_data_id = c.id
WHERE d.house_id = ?
`

type GetJoinedDevicesFromHouseRow struct {
	ID              int32
	HouseID         int32
	DeviceName      sql.NullString
	SetPoint        sql.NullFloat64
	Duration        sql.NullInt32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ClimateDataName string
	Unit            string
}

func (q *Queries) GetJoinedDevicesFromHouse(ctx context.Context, houseID int32) ([]GetJoinedDevicesFromHouseRow, error) {
	rows, err := q.db.QueryContext(ctx, getJoinedDevicesFromHouse, houseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetJoinedDevicesFromHouseRow
	for rows.Next() {
		var i GetJoinedDevicesFromHouseRow
		if err := rows.Scan(
			&i.ID,
			&i.HouseID,
			&i.DeviceName,
			&i.SetPoint,
			&i.Duration,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ClimateDataName,
			&i.Unit,
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
