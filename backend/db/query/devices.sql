-- name: CreateDevice :execlastid
INSERT INTO devices (climate_data_id, m304_id, name, rly) 
VALUES (?, ?, ?, ?);

-- name: GetDevicesFromM304 :many
SELECT 
    id, climate_data_id, m304_id, name, rly
FROM devices
WHERE m304_id = ?;

