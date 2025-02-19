CREATE TABLE devices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    climate_data_id INT NOT NULL,
    m304_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    rly INT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (climate_data_id) REFERENCES climate_datas(id) ON DELETE RESTRICT,
    FOREIGN KEY (m304_id) REFERENCES m304s(id) ON DELETE RESTRICT
);
