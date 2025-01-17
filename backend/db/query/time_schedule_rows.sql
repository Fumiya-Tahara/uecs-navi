-- name: CreateTimeSchedule :execlastid
INSERT INTO timeschedule_rows (time_schedule_id, start_time, end_time, workflow_id)
VALUES (?, ?, ?, ?);

-- name: GetTimeScheduleRowsFromTimeSchedule :many
SELECT id, time_schedule_id, start_time, end_time, workflow_id
FROM timeschedule_rows
WHERE time_schedule_id = ?;
