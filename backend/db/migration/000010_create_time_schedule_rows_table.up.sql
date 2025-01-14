CREATE TABLE time_schedules (
    id INT AUTO_INCREMENT PRIMARY KEY,
    time_schedule_id INT NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (time_schedule_id) REFERENCES time_schedules(id) ON DELETE RESTRICT
);
