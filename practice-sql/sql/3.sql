SELECT employees.id, employees.name, employees.surname
FROM employees, departments
WHERE employees.department_id = departments.id AND departments.title = 'IT';