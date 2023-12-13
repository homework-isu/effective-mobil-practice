CREATE TABLE employees (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    surname VARCHAR(255),
    department_id BIGINT,
    FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE CASCADE
);