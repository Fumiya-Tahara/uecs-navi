CREATE TABLE m304s (
    id INT AUTO_INCREMENT PRIMARY KEY,
    uecs_id VARCHAR(255) UNIQUE,
    house_id INT NOT NULL,
    mac_addr VARCHAR(255) NOT NULL,
    dhcp_flg BOOLEAN NOT NULL,
    ip_addr VARCHAR(255),
    net_mask VARCHAR(255),
    defgw VARCHAR(255),
    dns VARCHAR(255),
    vender_name VARCHAR(255) NOT NULL,
    node_name VARCHAR(255),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (house_id) REFERENCES houses(id) ON DELETE CASCADE
);
