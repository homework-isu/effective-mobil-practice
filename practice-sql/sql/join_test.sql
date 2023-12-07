SELECT 
    projects.title, 
    CONCAT(employees.name, ' ', employees.surname) as full_name,
    departments.title as department
FROM employees
LEFT JOIN employees_projects ON employees_projects.employee_id = employees.id
LEFT JOIN projects ON employees_projects.project_id = projects.id
LEFT JOIN departments ON employees.department_id = departments.id;


SELECT
    departments.title,
    CONCAT(employees.name, ' ', employees.surname) as full_name
FROM
    departments
FULL JOIN employees ON departments.id = employees.department_id;