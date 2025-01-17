-- name: CreateTimeSchedule :execlastid
INSERT INTO time_schedules (m304_id)
VALUES (?);

-- name: GetTimeScheduleFromM304 :one
SELECT id, m304_id
FROM time_schedules
WHERE m304_id = ?;

