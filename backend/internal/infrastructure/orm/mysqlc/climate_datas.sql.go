// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: climate_datas.sql

package mysqlc

import (
	"context"
)

const getAllClimateData = `-- name: GetAllClimateData :many
SELECT id, name, unit
FROM climate_datas
`

func (q *Queries) GetAllClimateData(ctx context.Context) ([]ClimateData, error) {
	rows, err := q.db.QueryContext(ctx, getAllClimateData)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClimateData
	for rows.Next() {
		var i ClimateData
		if err := rows.Scan(&i.ID, &i.Name, &i.Unit); err != nil {
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

const getClimateDataFromID = `-- name: GetClimateDataFromID :one
SELECT id, name, unit
FROM climate_datas
WHERE id = ?
`

func (q *Queries) GetClimateDataFromID(ctx context.Context, id int32) (ClimateData, error) {
	row := q.db.QueryRowContext(ctx, getClimateDataFromID, id)
	var i ClimateData
	err := row.Scan(&i.ID, &i.Name, &i.Unit)
	return i, err
}
