CREATE TABLE conditions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    climate_data_id INT NOT NULL,
    time_schedule_row_id INT NOT NULL,
    comparison_operator_id INT NOT NULL,
    set_point FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (climate_data_id) REFERENCES climate_datas(id) ON DELETE RESTRICT,
    FOREIGN KEY (time_schedule_row_id) REFERENCES time_schedule_rows(id) ON DELETE RESTRICT
);
