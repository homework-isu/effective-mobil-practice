-- SELECT departments.title as department, COUNT(employees.id)
-- FROM departments, employees
-- WHERE departments.id = employees.department_id
-- GROUP BY department;

SELECT departments.title as department, COUNT(employees.id)
FROM departments
JOIN employees ON departments.id = employees.department_id
GROUP BY department;