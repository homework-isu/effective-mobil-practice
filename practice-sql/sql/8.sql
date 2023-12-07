-- SELECT employees.id, CONCAT(employees.name, ' ', employees.surname) as full_name, departments.title as department
-- FROM employees
-- JOIN departments ON employees.department_id = departments.id;

SELECT employees.id, employees.name, employees.surname, departments.title as department
FROM employees
JOIN departments ON employees.department_id = departments.id;