CREATE TABLE conditions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    comparison_operator VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (climate_data_id) REFERENCES climate_datas(id) ON DELETE RESTRICT,
    FOREIGN KEY (time_schedule_row_id) REFERENCES time_schedule_rows(id) ON DELETE RESTRICT
);
