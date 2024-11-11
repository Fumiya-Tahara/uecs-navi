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
	M304ID        int32
	SensorID      int32
	DeviceName    sql.NullString
	Rly           sql.NullInt32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DeviceCondition struct {
	ID          int32
	DeviceID    int32
	OperationID int32
	Valid       bool
	SetPoint    sql.NullFloat64
	Duration    sql.NullInt32
	Operator    sql.NullInt32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type House struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type M304 struct {
	ID         int32
	UecsID     string
	MacAddr    string
	DhcpFlg    bool
	IpAddr     sql.NullString
	NetMask    sql.NullString
	Defgw      sql.NullString
	Dns        sql.NullString
	VenderName string
	NodeName   sql.NullString
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type M304Record struct {
	ID                int32
	M304ID            int32
	DeviceConditionID int32
	Block             string
	Valid             bool
	Position          int32
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Operation struct {
	ID        int32
	DeviceID  int32
	Name      string
	RlyOn     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Sensor struct {
	ID        int32
	CcmType   string
	Room      int32
	Region    int32
	Order     int32
	Priority  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TimeSchedule struct {
	ID                int32
	DeviceConditionID int32
	StartTime         string
	EndTime           string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
