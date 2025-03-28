-- name: CreateTimeScheduleRow :execlastid
INSERT INTO time_schedule_rows (time_schedule_id, start_time, end_time, workflow_id)
VALUES (?, ?, ?, ?);

-- name: GetTimeScheduleRowsFromTimeSchedule :many
SELECT id, time_schedule_id, start_time, end_time, workflow_id
FROM time_schedule_rows
WHERE time_schedule_id = ?;

-- name: UpdateTimeScheduleRow :exec
UPDATE time_schedule_rows
SET time_schedule_id = ?, start_time = ?, end_time = ?, workflow_id = ?
WHERE id = ?;
