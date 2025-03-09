// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: m304s.sql

package mysqlc

import (
	"context"
	"database/sql"
)

const createM304 = `-- name: CreateM304 :execlastid
INSERT INTO m304s (house_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateM304Params struct {
	HouseID    int32
	MacAddr    string
	DhcpFlg    bool
	IpAddr     sql.NullString
	NetMask    sql.NullString
	Defgw      sql.NullString
	Dns        sql.NullString
	VenderName string
	NodeName   sql.NullString
}

func (q *Queries) CreateM304(ctx context.Context, arg CreateM304Params) (int64, error) {
	result, err := q.db.ExecContext(ctx, createM304,
		arg.HouseID,
		arg.MacAddr,
		arg.DhcpFlg,
		arg.IpAddr,
		arg.NetMask,
		arg.Defgw,
		arg.Dns,
		arg.VenderName,
		arg.NodeName,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getAllM304s = `-- name: GetAllM304s :many
SELECT id, house_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name, created_at, updated_at FROM m304s
`

func (q *Queries) GetAllM304s(ctx context.Context) ([]M304, error) {
	rows, err := q.db.QueryContext(ctx, getAllM304s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []M304
	for rows.Next() {
		var i M304
		if err := rows.Scan(
			&i.ID,
			&i.HouseID,
			&i.MacAddr,
			&i.DhcpFlg,
			&i.IpAddr,
			&i.NetMask,
			&i.Defgw,
			&i.Dns,
			&i.VenderName,
			&i.NodeName,
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

const getM304FromID = `-- name: GetM304FromID :one
SELECT id, house_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name
FROM m304s
WHERE id = ?
`

type GetM304FromIDRow struct {
	ID         int32
	HouseID    int32
	MacAddr    string
	DhcpFlg    bool
	IpAddr     sql.NullString
	NetMask    sql.NullString
	Defgw      sql.NullString
	Dns        sql.NullString
	VenderName string
	NodeName   sql.NullString
}

func (q *Queries) GetM304FromID(ctx context.Context, id int32) (GetM304FromIDRow, error) {
	row := q.db.QueryRowContext(ctx, getM304FromID, id)
	var i GetM304FromIDRow
	err := row.Scan(
		&i.ID,
		&i.HouseID,
		&i.MacAddr,
		&i.DhcpFlg,
		&i.IpAddr,
		&i.NetMask,
		&i.Defgw,
		&i.Dns,
		&i.VenderName,
		&i.NodeName,
	)
	return i, err
}
