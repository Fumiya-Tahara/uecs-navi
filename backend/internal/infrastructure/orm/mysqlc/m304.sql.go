// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: m304.sql

package mysqlc

import (
	"context"
	"database/sql"
)

const createM304 = `-- name: CreateM304 :execlastid
INSERT INTO m304 (uecs_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name, ` + "`" + `rly_0` + "`" + `, ` + "`" + `rly_1` + "`" + `, ` + "`" + `rly_2` + "`" + `, ` + "`" + `rly_3` + "`" + `, ` + "`" + `rly_4` + "`" + `, ` + "`" + `rly_5` + "`" + `, ` + "`" + `rly_6` + "`" + `, ` + "`" + `rly_7` + "`" + `)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateM304Params struct {
	UecsID     string
	MacAddr    string
	DhcpFlg    bool
	IpAddr     sql.NullString
	NetMask    sql.NullString
	Defgw      sql.NullString
	Dns        sql.NullString
	VenderName string
	NodeName   sql.NullString
	Rly0       sql.NullInt32
	Rly1       sql.NullInt32
	Rly2       sql.NullInt32
	Rly3       sql.NullInt32
	Rly4       sql.NullInt32
	Rly5       sql.NullInt32
	Rly6       sql.NullInt32
	Rly7       sql.NullInt32
}

func (q *Queries) CreateM304(ctx context.Context, arg CreateM304Params) (int64, error) {
	result, err := q.db.ExecContext(ctx, createM304,
		arg.UecsID,
		arg.MacAddr,
		arg.DhcpFlg,
		arg.IpAddr,
		arg.NetMask,
		arg.Defgw,
		arg.Dns,
		arg.VenderName,
		arg.NodeName,
		arg.Rly0,
		arg.Rly1,
		arg.Rly2,
		arg.Rly3,
		arg.Rly4,
		arg.Rly5,
		arg.Rly6,
		arg.Rly7,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getM304FromUecsDevice = `-- name: GetM304FromUecsDevice :many
SELECT id, uecs_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name, ` + "`" + `rly_0` + "`" + `, ` + "`" + `rly_1` + "`" + `, ` + "`" + `rly_2` + "`" + `, ` + "`" + `rly_3` + "`" + `, ` + "`" + `rly_4` + "`" + `, ` + "`" + `rly_5` + "`" + `, ` + "`" + `rly_6` + "`" + `, ` + "`" + `rly_7` + "`" + `
FROM m304
WHERE ? IN (` + "`" + `rly_0` + "`" + `, ` + "`" + `rly_1` + "`" + `, ` + "`" + `rly_2` + "`" + `, ` + "`" + `rly_3` + "`" + `, ` + "`" + `rly_4` + "`" + `, ` + "`" + `rly_5` + "`" + `, ` + "`" + `rly_6` + "`" + `, ` + "`" + `rly_7` + "`" + `)
`

type GetM304FromUecsDeviceRow struct {
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
	Rly0       sql.NullInt32
	Rly1       sql.NullInt32
	Rly2       sql.NullInt32
	Rly3       sql.NullInt32
	Rly4       sql.NullInt32
	Rly5       sql.NullInt32
	Rly6       sql.NullInt32
	Rly7       sql.NullInt32
}

func (q *Queries) GetM304FromUecsDevice(ctx context.Context, rly0 sql.NullInt32) ([]GetM304FromUecsDeviceRow, error) {
	rows, err := q.db.QueryContext(ctx, getM304FromUecsDevice, rly0)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetM304FromUecsDeviceRow
	for rows.Next() {
		var i GetM304FromUecsDeviceRow
		if err := rows.Scan(
			&i.ID,
			&i.UecsID,
			&i.MacAddr,
			&i.DhcpFlg,
			&i.IpAddr,
			&i.NetMask,
			&i.Defgw,
			&i.Dns,
			&i.VenderName,
			&i.NodeName,
			&i.Rly0,
			&i.Rly1,
			&i.Rly2,
			&i.Rly3,
			&i.Rly4,
			&i.Rly5,
			&i.Rly6,
			&i.Rly7,
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
