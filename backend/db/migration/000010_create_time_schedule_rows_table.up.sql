CREATE TABLE time_schedule_rows (
    id INT AUTO_INCREMENT PRIMARY KEY,
    time_schedule_id INT NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    workflow_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (time_schedule_id) REFERENCES time_schedules(id) ON DELETE CASCADE,
    FOREIGN KEY (workflow_id) REFERENCES workflows(id) ON DELETE RESTRICT
);
