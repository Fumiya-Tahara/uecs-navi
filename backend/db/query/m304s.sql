-- name: CreateM304 :execlastid
INSERT INTO m304s (uecs_id, house_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetAllM304s :many
SELECT * FROM m304s;

-- name: GetM304FromID :one
SELECT id, uecs_id, house_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name
FROM m304s
WHERE id = ?;
