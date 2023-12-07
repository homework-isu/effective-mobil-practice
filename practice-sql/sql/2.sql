-- Добавление отделов
INSERT INTO departments (title) VALUES
    ('IT'),
    ('Marketing'),
    ('Finance');


-- Добавление сотрудников
INSERT INTO employees (name, surname, department_id) VALUES
    ('Danila', 'Ivashenko', 1),
    ('Fedor', 'Kuznetsov', 2),
    ('Egor', 'Egorov', 1),
    ('Ilia', 'Iliin', 3);


-- Добавление проектов
INSERT INTO projects (title) VALUES
    ('Project A'),
    ('Project B'),
    ('Project C');


-- Назначение сотрудников на проекты
INSERT INTO employees_projects (employee_id, project_id) VALUES
    (1, 1),
    (2, 1),
    (3, 2),
    (4, 3);
