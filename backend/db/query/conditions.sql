-- name: CreateCondition :execlastid
INSERT INTO conditions (climate_data_id, time_schedule_row_id, comparison_operator_id, set_point)
VALUES (?, ?, ?, ?);

-- name: GetConditionFromTimeScheduleRow :one
SELECT id, climate_data_id, time_schedule_row_id, comparison_operator_id, set_point
FROM conditions
WHERE time_schedule_row_id = ?;

-- name: UpdateCondition :exec
UPDATE conditions
SET climate_data_id = ?, time_schedule_row_id = ?, comparison_operator_id = ?, set_point = ?
WHERE id = ?;
