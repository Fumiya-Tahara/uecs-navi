// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package mysqlc

import (
	"database/sql"
	"time"
)

type ClimateData struct {
	ID   int32
	Name string
	Unit string
}

type Device struct {
	ID            int32
	HouseID       int32
	ClimateDataID int32
	DeviceName    sql.NullString
	SetPoint      sql.NullFloat64
	Duration      sql.NullInt32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type House struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
