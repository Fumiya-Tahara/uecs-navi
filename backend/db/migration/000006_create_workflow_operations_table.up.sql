CREATE TABLE workflow_operations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    workflow_id INT NOT NULL,
    relay_1 BOOLEAN,
    relay_2 BOOLEAN,
    relay_3 BOOLEAN,
    relay_4 BOOLEAN,
    relay_5 BOOLEAN,
    relay_6 BOOLEAN,
    relay_7 BOOLEAN,
    relay_8 BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (workflow_id) REFERENCES workflows(id) ON DELETE RESTRICT
);
